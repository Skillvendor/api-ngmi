package httpServer

import (
	"fmt"
	"os"

	"github.com/bmizerany/pat"

	"github.com/rs/cors"
)

func Run() {
	port := os.Getenv("PORT")
	env := os.Getenv("ENV")

	// mux := http.NewServeMux()

	mux := pat.New()

	InitRoutes(mux)

	allowedOrigins := []string{
		"http://localhost:3002",
		"http://localhost:3001",
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"https://ngmi-dev.netlify.app",
		"https://ngmilab.com",
		"https://ngmi-admin-dev.netlify.app",
	}

	if env == "prod" {
		allowedOrigins = []string{
			"https://ngmilab.com",
			"https://ngmi-admin.netlify.app",
			"https://stalwart-narwhal-8af26d.netlify.app",
		}
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		// Enable Debugging for testing, consider disabling in production
		Debug: env != "prod",
	})

	handler := c.Handler(mux)

	fmt.Println("Starting server from iPad")
	if env == "dev" {
		fmt.Println("Starting http")
		StartHttp(port, handler)
	} else {
		fmt.Println("Starting lambda")
		StartApexGateway(port, handler)
	}

}
