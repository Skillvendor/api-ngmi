package user

import (
	"api-ngmi/models"
	"api-ngmi/types"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding User"})
		w.WriteHeader(http.StatusInternalServerError)
	}

	if newUser.Address == "" {
		json.NewEncoder(w).Encode(types.StandardError{Message: "No Address Given"})
		w.WriteHeader(http.StatusBadRequest)
		return
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
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding User"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetPublicUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}

	address := r.URL.Query().Get(":address")

	newUser.Address = address

	if newUser.Address == "" {
		json.NewEncoder(w).Encode(types.StandardError{Message: "No Address provided"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser.Find()

	if newUser.Id == 0 {
		json.NewEncoder(w).Encode(types.StandardError{Message: "No User Found"})
		return
	}

	err := json.NewEncoder(w).Encode(newUser)

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding User"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
