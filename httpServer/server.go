package httpServer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Run() {
	port := os.Getenv("PORT")
	env := os.Getenv("ENV")

	mux := http.NewServeMux()

	handler := cors.Default().Handler(mux)

	InitRoutes(mux)

	fmt.Println("Starting server")
	if env == "dev" {
		fmt.Println("Starting http")
		StartHttp(port, handler)
	} else {
		fmt.Println("Starting lambda")
		StartApexGateway(port, handler)
	}

}
