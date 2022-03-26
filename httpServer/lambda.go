package httpServer

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
)

func StartApexGateway(port string, handler http.Handler) {
	log.Println("Starting lambda")
	log.Fatal(gateway.ListenAndServe(port, handler))
}
