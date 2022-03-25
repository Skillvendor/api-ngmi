package httpServer

import (
	"log"
	"net/http"
)

func StartHttp(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
