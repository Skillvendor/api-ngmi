package httpServer

import (
	"log"

	"github.com/apex/gateway"
)

func StartApexGateway(port string) {
	log.Fatal(gateway.ListenAndServe(port, nil))
}
