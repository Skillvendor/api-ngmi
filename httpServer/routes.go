package httpServer

import (
	"go-rarity/api/collection"
	"go-rarity/api/s3"
	"go-rarity/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	// public
	mux.HandleFunc("/api/collection/all", collection.GetPublishedCollections)
	mux.HandleFunc("/api/collection/show", collection.GetCollection)

	// private
	mux.HandleFunc("/api/collection/create", middleware.CheckApiKey(collection.CreateCollection))
	mux.HandleFunc("/api/collection/update", middleware.CheckApiKey(collection.UpdateCollection))
	mux.HandleFunc("/api/collection/delete", middleware.CheckApiKey(collection.DeleteCollection))
	mux.HandleFunc("/api/collection/approveReview", middleware.CheckApiKey(collection.ApproveReview))
	mux.HandleFunc("/api/collection/admin", middleware.CheckApiKey(collection.GetAllCollections))

	mux.HandleFunc("/api/signedUrl", middleware.CheckApiKey(s3.GetSignedUploadUrl))
}
