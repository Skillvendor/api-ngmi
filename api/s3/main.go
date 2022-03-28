package s3

import (
	"encoding/json"
	"go-rarity/s3service"
	"log"
	"net/http"
)

func GetSignedUrl(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(s3service.GetSignedUrl())

	if err != nil {
		log.Println("ERROR Encoding S3 object", err)
	}
}
