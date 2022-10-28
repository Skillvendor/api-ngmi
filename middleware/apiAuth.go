package middleware

import (
	"api-ngmi/models"
	"api-ngmi/services/auth"
	userService "api-ngmi/services/user"
	"api-ngmi/types"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func CheckApiKey(handler func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		apiKey := os.Getenv("RARITY_API_KEY")
		ah := r.Header.Get("Authorization")

		if ah == apiKey {
			handler(w, r)
		} else {
			return &types.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        errors.New("unauthorized"),
			}
		}

		return nil
	}
}

func CheckJWTToken(handler func(w http.ResponseWriter, r *http.Request) error, accessLevel int) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		authHeader := r.Header.Get("Authorization")

		claims, tkn, err := auth.DecodeJWT(authHeader)

		if !tkn.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					return &types.RequestError{
						StatusCode: http.StatusUnauthorized,
						Err:        errors.New("token malformed"),
					}
				}

				if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					return &types.RequestError{
						StatusCode: http.StatusUnauthorized,
						Err:        errors.New("token Expired"),
					}
				}

				return &types.RequestError{
					StatusCode: http.StatusBadRequest,
					Err:        errors.New("did you provide a token?"),
				}
			}

			return &types.RequestError{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("not a valid token, did you even provide one?"),
			}
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return &types.RequestError{
					StatusCode: http.StatusUnauthorized,
					Err:        errors.New("unauthorized"),
				}
			}
		}

		user := models.User{Address: claims.Address}
		user.Find()

		if user.AuthToken != authHeader {
			return &types.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        errors.New("address should have a different token"),
			}
		}

		if userService.AccessLevelFor(claims.Address) < accessLevel {
			return &types.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        errors.New("no access to this resource"),
			}
		}

		ctx := context.WithValue(r.Context(), "user", user)
		handler(w, r.WithContext(ctx))

		return nil
	}
}

func CheckAdminJWTToken(handler func(w http.ResponseWriter, r *http.Request) error, accessLevel int) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		authHeader := r.Header.Get("Authorization")

		claims, tkn, err := auth.DecodeJWT(authHeader)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return &types.RequestError{
					StatusCode: http.StatusUnauthorized,
					Err:        errors.New("unauthorized"),
				}
			}

			return &types.RequestError{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("did you provide a token?"),
			}
		}
		if !tkn.Valid {
			return &types.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        errors.New("unauthorized"),
			}
		}

		admin := models.Admin{Username: claims.Username}
		admin.Find()

		if admin.AuthToken != authHeader {
			return &types.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        errors.New("address should have a different token"),
			}
		}

		if admin.AccessLevel < accessLevel {
			return &types.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        errors.New("no access to this resource"),
			}
		}

		ctx := context.WithValue(r.Context(), "user", admin)
		handler(w, r.WithContext(ctx))

		fmt.Println("this is the admin", admin)

		return nil
	}
}
