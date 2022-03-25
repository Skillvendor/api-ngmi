package httpServer

import (
	"log"

	"github.com/apex/gateway/v2"
)

func StartApexGateway(port string) {
	log.Fatal(gateway.ListenAndServe(port, nil))
}
