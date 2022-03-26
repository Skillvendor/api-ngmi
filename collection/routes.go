package collection

import (
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/hello", MyFunc)
	mux.HandleFunc("/api/collection/create", CreateCollection)
	mux.HandleFunc("/api/collection/all", GetAllCollections)
}
