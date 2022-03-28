package collection

import (
	"encoding/json"
	"go-rarity/models"
	"log"
	"net/http"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			log.Println("Error encoding collection", err)
		}
		myCollection.Published = true

		myCollection.Save()

		err = json.NewEncoder(w).Encode(myCollection)

		if err != nil {
			log.Println("ERROR Encoding", err)
		}
	default:
		w.Write([]byte("Come on man... it's a post"))
	}
}

func GetCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			log.Println("Error encoding collection", err)
		}

		myCollection.Find()

		err = json.NewEncoder(w).Encode(myCollection)

		if err != nil {
			log.Println("ERROR Encoding", err)
		}
	default:
		w.Write([]byte("Come on man... it's a post"))
	}
}

func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			log.Println("Error encoding collection", err)
		}

		json.NewEncoder(w).Encode(models.DeleteCollectionById(int(myCollection.Id)))
	default:
		w.Write([]byte("Come on man... delete method and an id"))
	}
}

func ApproveReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			log.Println("Error encoding collection", err)
		}

		json.NewEncoder(w).Encode(models.ApproveReview(int(myCollection.Id)))
	default:
		w.Write([]byte("Come on man... post method and an id"))
	}
}

func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAllCollections())
}
