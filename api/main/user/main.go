package user

import (
	"api-ngmi/models"
	"api-ngmi/redis"
	"api-ngmi/services/auth"
	userService "api-ngmi/services/user"
	"api-ngmi/types"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) error {
	newUser := models.User{}

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error decoding user"),
		}
	}

	if newUser.Address == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no address given"),
		}
	}

	newUser.Find()

	if newUser.Id == 0 {
		// user doesn't exist so we try creating
		rand.Seed(time.Now().UnixNano())
		newUser.Nonce = strconv.Itoa(rand.Int())
		newUser.AccessLevel = 0
		newUser.AuthToken = ""
		newUser.Save()
	} else {
		newUser.AccessLevel = userService.AccessLevelFor(newUser.Address)
	}

	err = json.NewEncoder(w).Encode(newUser)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding user"),
		}
	}

	return nil
}

type LoginData struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func Update(w http.ResponseWriter, r *http.Request) error {
	var user models.User = r.Context().Value("user").(models.User)

	newUser := LoginData{}

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error decoding user"),
		}
	}

	shouldSave := false
	if newUser.Username != "" {
		user.Username = newUser.Username
		shouldSave = true
	}

	if newUser.Password != "" {
		hashedPass, _ := auth.HashPassword(newUser.Password)
		user.Password = hashedPass
		shouldSave = true
	}

	if shouldSave {
		saved := user.UpdateEmailPass()
		if !saved {
			return &types.RequestError{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("can't save user"),
			}
		}
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding user"),
		}
	}

	return nil
}

func Show(w http.ResponseWriter, r *http.Request) error {
	// address := r.URL.Query().Get(":address")

	var user models.User = r.Context().Value("user").(models.User)

	if user.Address == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no address given"),
		}
	}

	user.Find()

	if user.Id == 0 {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no user found"),
		}
	}

	user.AccessLevel = userService.AccessLevelFor(user.Address)

	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error encoding user"),
		}
	}

	return nil
}

func ResetAccessLevelCache(w http.ResponseWriter, r *http.Request) error {
	var user models.User = r.Context().Value("user").(models.User)

	_, err := redis.PurgeAccessLevelCacheDBFor(user.Address)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't purge cache for user"),
		}
	}

	// empty
	newUser := models.User{}
	err = json.NewEncoder(w).Encode(newUser)
	w.WriteHeader(http.StatusOK)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error encoding user"),
		}
	}

	return nil
}
