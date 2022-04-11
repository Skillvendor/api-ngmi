package s3

import (
	"encoding/json"
	"go-rarity/services/s3"
	"log"
	"net/http"
)

func GetSignedUploadUrl(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(s3.GetSignedUploadUrl())

	if err != nil {
		log.Println("ERROR Encoding S3 object", err)
	}
}
