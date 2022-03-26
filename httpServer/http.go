package httpServer

import (
	"log"
	"net/http"
)

func StartHttp(port string, handler http.Handler) {
	log.Fatal(http.ListenAndServe(port, handler))
}
