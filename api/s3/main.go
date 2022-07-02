package s3

import (
	"api-ngmi/services/s3"
	"api-ngmi/types"
	"encoding/json"
	"net/http"
)

func GetSignedUploadUrlAssets(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(s3.AssetsBucket.GetSignedUploadUrl())

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 Upload Url"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetSignedUploadUrlReports(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(s3.ReportsBucket.GetSignedUploadUrl())

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 Upload Url"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type SignedReadUrlResp struct {
	Url string
	Key string
}

func GetSignedDownloadUrlReports(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get(":key")

	if key == "" {
		json.NewEncoder(w).Encode(types.StandardError{Message: "No key provided"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := s3.ReportsBucket.GetSignedReadUrl(key)

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Can't get url"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(SignedReadUrlResp{Url: url, Key: key})

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 Report Download Url"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
