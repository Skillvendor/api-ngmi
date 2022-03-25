package animal

import (
	"net/http"
)

func InitAnimalRoutes() {
	http.HandleFunc("/", MyFunc)
	http.HandleFunc("/api/hello", MyFunc)
}
