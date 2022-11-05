package middleware

import (
	"api-ngmi/types"
	"encoding/json"
	"net/http"
)

func LimitRequestLength(handler func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

		return handler(w, r)
	}
}

func ErrorHandler(handler func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			re, ok := err.(*types.RequestError)
			if ok {
				w.WriteHeader(re.StatusCode)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(&types.StandardError{Message: err.Error()})
		}
	}
}

func ApplyStandardMiddlewares(handler func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return ErrorHandler(LimitRequestLength(handler))
}
