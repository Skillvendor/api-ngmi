package httpServer

import (
	"fmt"
	"go-rarity/collection"
	"os"
)

func Run() {
	port := os.Getenv("PORT")
	env := os.Getenv("ENV")

	collection.InitAnimalRoutes()

	fmt.Println("Starting server")
	if env == "dev" {
		fmt.Println("Starting http")
		StartHttp(port)
	} else {
		fmt.Println("Starting lambda")
		StartApexGateway(port)
	}

}
