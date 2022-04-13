package user

import (
	"encoding/json"
	"go-rarity/models"
	"math/rand"
	"time"

	"log"
	"net/http"

	"github.com/jinzhu/copier"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newUser := models.User{}

		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			log.Println("Error encoding collection", err)
		}

		if newUser.Address == "" {
			w.Write([]byte("Come on man... give me an address"))
			return
		}

		newUser.Find()

		if newUser.Id == 0 {
			// user doesn't exist so we try creating
			rand.Seed(time.Now().UnixNano())
			newUser.Nonce = rand.Int()
			newUser.AccessLevel = 1
			newUser.AuthToken = ""
			newUser.Save()
		}

		err = json.NewEncoder(w).Encode(newUser)

		if err != nil {
			log.Println("ERROR Encoding", err)
		}
	default:
		w.Write([]byte("Come on man... it's a post"))
	}
}

func GetPublicUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newUser := models.User{}

		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			log.Println("Error encoding collection", err)
			return
		}

		if newUser.Address == "" {
			w.Write([]byte("Come on man... give me an address"))
			return
		}

		newUser.Find()

		if newUser.Id == 0 {
			// user doesn't exist so we try creating
			w.Write([]byte("Can't find user"))
			return
		}

		publicUser := models.PublicUser{}
		copier.Copy(&publicUser, &newUser)

		err = json.NewEncoder(w).Encode(publicUser)

		if err != nil {
			log.Println("ERROR Encoding", err)
		}
	default:
		w.Write([]byte("Please use a post"))
	}
}
