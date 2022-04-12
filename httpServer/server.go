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

	InitRoutes(mux)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3001",
			"https://rarity-admin-dev.netlify.app",
			"https://rarity-admin-prod.netlify.app",
			"https://magic-carpet-dev.netlify.app",
			"https://magic-carpet.netlify.app",
			"http://local.magic-carpet-admin.com:3001",
			"https://local.magic-carpet-admin.com:3001",
			"https://ngmilab.com",
		},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		// Enable Debugging for testing, consider disabling in production
		Debug: env != "prod",
	})

	handler := c.Handler(mux)

	fmt.Println("Starting server")
	if env == "dev" {
		fmt.Println("Starting http")
		StartHttp(port, handler)
	} else {
		fmt.Println("Starting lambda")
		StartApexGateway(port, handler)
	}

}
