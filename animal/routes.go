package animal

import (
	"net/http"
)

func InitAnimalRoutes() {
	http.HandleFunc("/api/hello", MyFunc)
	http.HandleFunc("/api/hello2", MyFunc2)
}
