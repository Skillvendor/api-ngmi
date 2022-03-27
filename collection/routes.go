package collection

import (
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/collection/create", CreateCollection)
	mux.HandleFunc("/api/collection/all", GetAllCollections)
	mux.HandleFunc("/api/collection/delete", DeleteCollection)
	mux.HandleFunc("/api/collection/approveReview", ApproveReview)
}
