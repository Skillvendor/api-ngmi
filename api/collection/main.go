package collection

import (
	"encoding/json"
	"go-rarity/models"
	collectionService "go-rarity/services/collection"
	"go-rarity/types"

	"net/http"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Can't decode collection"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		myCollection.Published = false

		myCollection.Save()

		err = json.NewEncoder(w).Encode(myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Can't encode collection"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Make a post request"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PATCH", "PUT":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding collection"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		myCollection.Update()

		err = json.NewEncoder(w).Encode(myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding Token"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a Patch or Put"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding Token"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		myCollection.Find()

		err = json.NewEncoder(w).Encode(collectionService.TransformToS3Urls(myCollection))

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 url transformations"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a post"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding Collection"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(models.DeleteCollectionById(int(myCollection.Id)))
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a delete call"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ApproveReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		myCollection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&myCollection)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding Token"})
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(models.ApproveReview(int(myCollection.Id)))
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a post"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetPublishedCollections(w http.ResponseWriter, r *http.Request) {
	collections := collectionService.GetPublishedCollections()

	json.NewEncoder(w).Encode(collectionService.MapToS3Urls(collections))
}

func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	collections := models.GetAllCollections()

	json.NewEncoder(w).Encode(collectionService.MapToS3Urls(collections))
}
