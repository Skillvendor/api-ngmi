package user

import (
	"api-ngmi/models"
	"api-ngmi/redis"
	userService "api-ngmi/services/user"
	"api-ngmi/types"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"net/http"
)

type UserData struct {
	User        models.User
	AccessLevel int
}

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
		newUser.AccessLevel = 1
		newUser.AuthToken = ""
		newUser.Save()
	}

	err = json.NewEncoder(w).Encode(UserData{User: newUser, AccessLevel: 0})

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

	err := json.NewEncoder(w).Encode(UserData{User: user, AccessLevel: userService.AccessLevelFor(user.Address)})

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

	_, err := redis.PurgeAccessLevelCacheFor(user.Address)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't purge cache for user"),
		}
	}

	return nil
}
