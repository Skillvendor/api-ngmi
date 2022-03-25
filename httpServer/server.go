package httpServer

import (
	"flag"
	"fmt"
	"go-rarity/animal"
	"os"
)

func Run() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	env := os.Getenv("ENV")

	if *port == -1 {
		panic("SPECIFY PORT")
	}
	portStr := fmt.Sprintf(":%d", *port)

	animal.InitAnimalRoutes()

	fmt.Println("Starting server")
	if env == "dev" {
		fmt.Println("Starting http")
		StartHttp(portStr)
	} else {
		fmt.Println("Starting lambda")
		StartApexGateway(portStr)
	}

}
