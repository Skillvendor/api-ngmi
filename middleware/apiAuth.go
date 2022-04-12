package middleware

import (
	"context"
	"go-rarity/models"
	"go-rarity/services/auth"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func CheckApiKey(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("RARITY_API_KEY")
		ah := r.Header.Get("Authorization")

		if ah == apiKey {
			handler(w, r)
		} else {
			w.Write([]byte("Invalid Key Hacker!"))
		}
	}
}

func CheckJWTToken(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		claims, tkn, err := auth.DecodeJWT(authHeader)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user := models.User{Address: claims.Address}
		user.Find()

		if user.AuthToken != authHeader {
			w.Write([]byte("Why do you steal tokens ?"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		handler(w, r.WithContext(ctx))
	}
}
