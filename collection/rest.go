package collection

import (
	"encoding/json"
	"go-rarity/models"
	"log"
	"net/http"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	log.Println("HERE")
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			log.Println("Error encoding collection", err)
		}

		log.Println("ENCODED", myCollection)

		myCollection.Save()

		err = json.NewEncoder(w).Encode(myCollection)

		if err != nil {
			log.Println("ERROR Encoding", err)
		}
	default:
		w.Write([]byte("Come on man... it's a post"))
	}
}

func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAllCollections())
}

func MyFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO"))
}
