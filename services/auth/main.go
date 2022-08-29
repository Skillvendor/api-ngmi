package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JWTPayload struct {
	Address        string
	Username       string
	ExpirationTime time.Time
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var AuthTokenExpirationTime = time.Now().Add(24 * time.Hour)

func CreateJWT(payload JWTPayload) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := payload.ExpirationTime

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Address:  payload.Address,
		Username: payload.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", errors.New("can't generate token")
	}

	return tokenString, nil
}

func DecodeJWT(token string) (*Claims, *jwt.Token, error) {
	jwtKey := os.Getenv("JWT_SECRET")

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	return claims, tkn, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
