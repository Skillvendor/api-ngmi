package auth

import (
	"api-ngmi/models"
	"api-ngmi/services/auth"
	userService "api-ngmi/services/user"
	"api-ngmi/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt"
)

type SignatureData struct {
	Address   string `json:"address,omitempty"`
	Signature string `json:"signature,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token,omitempty"`
}

func verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return strings.EqualFold(from, recoveredAddr.Hex())
}

func Authentication(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("I AM AUTHENTICATING")
	newSignature := SignatureData{}

	err := json.NewDecoder(r.Body).Decode(&newSignature)
	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error decoding body"),
		}
	}

	if newSignature.Address == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no address provided"),
		}
	}

	user := models.User{Address: newSignature.Address}
	found := user.Find()

	if !found {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("user does not exist"),
		}
	}

	msg := "I am signing my one-time nonce: " + user.Nonce

	signed := verifySig(user.Address, newSignature.Signature, []byte(msg))

	if !signed {
		return &types.RequestError{
			StatusCode: http.StatusUnauthorized,
			Err:        errors.New("signature invalid"),
		}
	}

	userAccessLevel := userService.AccessLevelFor(user.Address)

	newPayload := auth.JWTPayload{
		Address:        user.Address,
		AccessLevel:    userAccessLevel,
		ExpirationTime: auth.AuthTokenExpirationTime(),
	}

	authToken, jwtError := auth.CreateJWT(newPayload)

	fmt.Println("I CREATED A NEW AUTH TOKEN", authToken)

	if jwtError != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't generate token"),
		}
	}

	user.AuthToken = authToken
	rand.Seed(time.Now().UnixNano())
	user.Nonce = strconv.Itoa(rand.Int())
	user.Update()

	err = json.NewEncoder(w).Encode(AuthResponse{Token: authToken})

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding token"),
		}
	}

	// Finally, we set the client cookie for "Authorization" as the JWT we just generated
	// we also set an expiry time
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   authToken,
		Expires: auth.AuthTokenExpirationTime(),
	})

	return nil
}

func TestJWT(w http.ResponseWriter, r *http.Request) error {
	var newUser models.Admin = r.Context().Value("user").(models.Admin)
	log.Println("I actually got the user", newUser.Username)
	authHeader := r.Header.Get("Authorization")

	claims, tkn, err := auth.DecodeJWT(authHeader)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return nil
		}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	user := models.User{Address: claims.Address}
	user.Find()

	ctx := context.WithValue(r.Context(), "user", user)
	log.Println(ctx)
	w.Write([]byte("Seems fine"))

	return nil
}
