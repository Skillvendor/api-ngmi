package middleware

import (
	"net/http"
	"os"
)

func CheckApiKey(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("RARITY_API_KEY")
		ah := r.Header.Get("Authorization")

		if ah == apiKey {
			handler(w, r)
		} else {
			w.Write([]byte("Invalid Key Hacker!"))
		}
	}
}
