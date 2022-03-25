package httpServer

import (
	"log"

	"github.com/apex/gateway"
)

func StartApexGateway(port string) {
	log.Println("Starting lambda")
	log.Fatal(gateway.ListenAndServe(port, nil))
}
