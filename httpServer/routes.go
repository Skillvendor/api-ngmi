package httpServer

import (
	"go-rarity/api/collection"
	"go-rarity/api/s3"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/collection/create", collection.CreateCollection)
	mux.HandleFunc("/api/collection/all", collection.GetPublishedCollections)
	mux.HandleFunc("/api/collection/delete", collection.DeleteCollection)
	mux.HandleFunc("/api/collection/approveReview", collection.ApproveReview)
	mux.HandleFunc("/api/collection/show", collection.GetCollection)
	mux.HandleFunc("/api/collection/admin", collection.GetAllCollections)

	mux.HandleFunc("/api/signedUrl", s3.GetSignedUrl)
}
