package auth

import (
	"context"
	"encoding/json"
	"go-rarity/models"
	"go-rarity/services/auth"
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

func CreateJWT(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		address := "dawfdeaiudjwiu32132131321321"
		newPayload := auth.JWTPayload{
			Address:        address,
			ExpirationTime: auth.AuthTokenExpirationTime,
		}

		authToken := auth.CreateJWT(newPayload)

		w.Write([]byte(authToken))
	default:
		w.Write([]byte("Come on man... it's a post"))
	}
}

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

	log.Println("this is verify sig ", msg, sigHex, recoveredAddr, from)

	return from == recoveredAddr.Hex()
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

		signed := verifySig(strings.ToLower(user.Address), newSignature.Signature, []byte(msg))

		if !signed {
			log.Println("this is msg ", msg)
			w.Write([]byte("Are you trying to hack?"))
			return
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
			log.Println("ERROR Encoding Token", err)
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
		w.Write([]byte("Come on man... it's a post"))
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
