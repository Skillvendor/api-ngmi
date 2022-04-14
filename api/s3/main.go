package s3

import (
	"encoding/json"
	"go-rarity/services/s3"
	"go-rarity/types"
	"net/http"
)

func GetSignedUploadUrl(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(s3.GetSignedUploadUrl())

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 Upload Url"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
