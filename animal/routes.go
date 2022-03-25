package animal

import (
	"net/http"
)

func InitAnimalRoutes() {
	http.HandleFunc("/", MyFunc2)
	http.HandleFunc("/api/hello", MyFunc)
	http.HandleFunc("/api/hello2", MyFunc2)
}
