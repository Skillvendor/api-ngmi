package user

import (
	"encoding/json"
	"api-ngmi/models"
	"api-ngmi/types"
	"math/rand"
	"strconv"
	"time"

	"net/http"

	"github.com/jinzhu/copier"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
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
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Post method"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetPublicUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newUser := models.User{}

		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding User"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newUser.Address == "" {
			json.NewEncoder(w).Encode(types.StandardError{Message: "No Address provided"})
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newUser.Find()

		if newUser.Id == 0 {
			// user doesn't exist so we try creating
			json.NewEncoder(w).Encode(types.StandardError{Message: "No User Found"})
			return
		}

		publicUser := models.PublicUser{}
		copier.Copy(&publicUser, &newUser)

		err = json.NewEncoder(w).Encode(publicUser)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding User"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "It should be a POST"})
		w.WriteHeader(http.StatusBadRequest)
	}
}
