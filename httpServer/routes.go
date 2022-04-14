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
	mux.HandleFunc("/api/auth/test", middleware.CheckJWTToken(auth.TestJWT, 1))

	// private

	// collection
	mux.HandleFunc("/api/collection/create", middleware.CheckJWTToken(collection.CreateCollection, 5))
	mux.HandleFunc("/api/collection/update", middleware.CheckJWTToken(collection.UpdateCollection, 5))
	mux.HandleFunc("/api/collection/delete", middleware.CheckJWTToken(collection.DeleteCollection, 6))
	mux.HandleFunc("/api/collection/approveReview", middleware.CheckJWTToken(collection.ApproveReview, 6))
	mux.HandleFunc("/api/collection/admin", middleware.CheckJWTToken(collection.GetAllCollections, 5))

	// s3
	mux.HandleFunc("/api/signedUrl", middleware.CheckJWTToken(s3.GetSignedUploadUrl, 5))
}
