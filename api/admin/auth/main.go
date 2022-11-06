package admin

import (
	"api-ngmi/models"
	"api-ngmi/services/auth"
	"api-ngmi/types"
	"encoding/json"
	"errors"

	"net/http"
)

type LoginInfo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token,omitempty"`
}

func Authentication(w http.ResponseWriter, r *http.Request) error {
	newLoginInfo := LoginInfo{}

	err := json.NewDecoder(r.Body).Decode(&newLoginInfo)
	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error decoding body"),
		}
	}

	if newLoginInfo.Username == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no username provided"),
		}
	}

	if newLoginInfo.Password == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no password provided"),
		}
	}

	admin := models.Admin{Username: newLoginInfo.Username}
	found := admin.Find()

	if !found {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("admin does not exist"),
		}
	}

	okPass := auth.CheckPasswordHash(newLoginInfo.Password, admin.Password)

	if !okPass {
		return &types.RequestError{
			StatusCode: http.StatusUnauthorized,
			Err:        errors.New("password invalid"),
		}
	}

	newPayload := auth.JWTPayload{
		Username:       admin.Username,
		ExpirationTime: auth.AuthTokenExpirationTime(),
	}

	authToken, jwtError := auth.CreateJWT(newPayload)

	if jwtError != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't generate token"),
		}
	}

	admin.AuthToken = authToken
	updated := admin.Update()

	if !updated {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("something went wrong"),
		}
	}

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
