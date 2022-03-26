package collection

import (
	"net/http"
)

func InitAnimalRoutes() {
	http.HandleFunc("/api/hello", MyFunc)
	http.HandleFunc("/api/collection/create", CreateCollection)
	http.HandleFunc("/api/collection/all", GetAllCollections)
}
