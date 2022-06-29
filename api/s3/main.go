package s3

import (
	"api-ngmi/services/s3"
	"api-ngmi/types"
	"encoding/json"
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
