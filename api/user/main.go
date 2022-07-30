package user

import (
	"api-ngmi/models"
	"api-ngmi/types"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) error {
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

	err = json.NewEncoder(w).Encode(newUser)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding user"),
		}
	}

	return nil
}

func GetPublicUser(w http.ResponseWriter, r *http.Request) error {
	newUser := models.User{}

	address := r.URL.Query().Get(":address")

	newUser.Address = address

	if newUser.Address == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no address given"),
		}
	}

	newUser.Find()

	if newUser.Id == 0 {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("no user found"),
		}
	}

	err := json.NewEncoder(w).Encode(newUser)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error encoding user"),
		}
	}

	return nil
}
