package s3

import (
	"api-ngmi/services/s3"
	"api-ngmi/types"
	"encoding/json"
	"errors"
	"net/http"
)

func GetSignedUploadUrlAssets(w http.ResponseWriter, r *http.Request) error {
	err := json.NewEncoder(w).Encode(s3.AssetsBucket.GetSignedUploadUrl())

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Error encoding s3 upload url"),
		}
	}

	return nil
}

func GetSignedUploadUrlReports(w http.ResponseWriter, r *http.Request) error {
	err := json.NewEncoder(w).Encode(s3.ReportsBucket.GetSignedUploadUrl())

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("Error encoding s3 upload url"),
		}
	}

	return nil
}

type SignedReadUrlResp struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func GetSignedDownloadUrlReports(w http.ResponseWriter, r *http.Request) error {
	key := r.URL.Query().Get(":key")

	if key == "" {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("No Key provided"),
		}
	}

	url, err := s3.ReportsBucket.GetSignedReadUrl(key)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't get url"),
		}
	}

	err = json.NewEncoder(w).Encode(SignedReadUrlResp{Url: url, Key: key})

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error Encoding S3 Report Download Url"),
		}
	}

	return nil
}
