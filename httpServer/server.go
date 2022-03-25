package httpServer

import (
	"flag"
	"fmt"
	"go-rarity/animal"
	"os"
)

func Run() {
	port := os.Getenv("PORT")
	env := os.Getenv("ENV")
	flag.Parse()

	animal.InitAnimalRoutes()

	fmt.Println("Starting server")
	if env == "dev" {
		fmt.Println("Starting http")
		StartHttp(port)
	} else {
		fmt.Println("Starting lambda")
		StartApexGateway(port)
	}

}
