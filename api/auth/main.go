package auth

import (
	"api-ngmi/models"
	"api-ngmi/services/auth"
	"api-ngmi/types"
	"context"
	"encoding/json"
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
	Address   string
	Signature string
}

type AuthResponse struct {
	Token string
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

	return strings.ToLower(from) == strings.ToLower(recoveredAddr.Hex())
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newSignature := SignatureData{}

		err := json.NewDecoder(r.Body).Decode(&newSignature)
		if err != nil {
			log.Println("Error decoding body", err)
			return
		}

		if newSignature.Address == "" {
			log.Println("No Address provided")
			return
		}

		user := models.User{Address: newSignature.Address}
		user.Find()

		msg := "I am signing my one-time nonce: " + user.Nonce

		signed := verifySig(user.Address, newSignature.Signature, []byte(msg))

		if !signed {
			json.NewEncoder(w).Encode(types.StandardError{Message: "No Hackers Allowed"})
			w.WriteHeader(http.StatusUnauthorized)
		}

		address := user.Address
		newPayload := auth.JWTPayload{
			Address:        address,
			ExpirationTime: auth.AuthTokenExpirationTime,
		}

		authToken := auth.CreateJWT(newPayload)

		user.AuthToken = authToken
		rand.Seed(time.Now().UnixNano())
		user.Nonce = strconv.Itoa(rand.Int())
		user.Update()

		err = json.NewEncoder(w).Encode(AuthResponse{Token: authToken})

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding Token"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Finally, we set the client cookie for "Authorization" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "Authorization",
			Value:   authToken,
			Expires: auth.AuthTokenExpirationTime,
		})
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a POST request"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TestJWT(w http.ResponseWriter, r *http.Request) {
	anotherUser := r.Context().Value("user")
	log.Println("I actually got the user", anotherUser)
	authHeader := r.Header.Get("Authorization")

	claims, tkn, err := auth.DecodeJWT(authHeader)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := models.User{Address: claims.Address}
	user.Find()

	ctx := context.WithValue(r.Context(), "user", user)
	log.Println(ctx)
	w.Write([]byte("Seems fine"))
}
