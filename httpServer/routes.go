package httpServer

import (
	"go-rarity/api/auth"
	"go-rarity/api/collection"
	"go-rarity/api/s3"
	"go-rarity/api/user"
	"go-rarity/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	// public

	// collection
	mux.HandleFunc("/api/collection/all", collection.GetPublishedCollections)
	mux.HandleFunc("/api/collection/show", collection.GetCollection)

	// user
	mux.HandleFunc("/api/user/create", user.CreateUser)
	mux.HandleFunc("/api/user/show", user.GetPublicUser)

	// auth
	mux.HandleFunc("/api/auth/authentication", auth.Authentication)
	mux.HandleFunc("/api/auth/test", middleware.CheckJWTToken(auth.TestJWT))

	// private

	// collection
	mux.HandleFunc("/api/collection/create", middleware.CheckApiKey(collection.CreateCollection))
	mux.HandleFunc("/api/collection/update", middleware.CheckApiKey(collection.UpdateCollection))
	mux.HandleFunc("/api/collection/delete", middleware.CheckApiKey(collection.DeleteCollection))
	mux.HandleFunc("/api/collection/approveReview", middleware.CheckApiKey(collection.ApproveReview))
	mux.HandleFunc("/api/collection/admin", middleware.CheckApiKey(collection.GetAllCollections))

	// s3
	mux.HandleFunc("/api/signedUrl", middleware.CheckApiKey(s3.GetSignedUploadUrl))
}
