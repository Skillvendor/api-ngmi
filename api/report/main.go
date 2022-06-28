package report

import (
	"encoding/json"
	"api-ngmi/models"
	reportService "api-ngmi/services/report"
	"api-ngmi/types"

	"net/http"
)

func CreateReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newReport := models.Report{}

		err := json.NewDecoder(r.Body).Decode(&newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Can't decode report"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newReport.Published = false

		newReport.Save()

		err = json.NewEncoder(w).Encode(newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Can't encode report"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Make a post request"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PATCH", "PUT":
		newReport := models.Report{}

		err := json.NewDecoder(r.Body).Decode(&newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding report"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		newReport.Update()

		err = json.NewEncoder(w).Encode(newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding Token"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a Patch or Put"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newReport := models.Report{}

		err := json.NewDecoder(r.Body).Decode(&newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding Token"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		newReport.Find()

		err = json.NewEncoder(w).Encode(reportService.TransformToS3Urls(newReport))

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 url transformations"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a post"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func DeleteReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newReport := models.Report{}

		err := json.NewDecoder(r.Body).Decode(&newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding report"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(models.DeleteReportById(int(newReport.Id)))
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a delete call"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ApproveReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newReport := models.Report{}

		err := json.NewDecoder(r.Body).Decode(&newReport)

		if err != nil {
			json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding Token"})
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(models.ApproveReview(int(newReport.Id)))
	default:
		json.NewEncoder(w).Encode(types.StandardError{Message: "Expected a post"})
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetPublishedReports(w http.ResponseWriter, r *http.Request) {
	reports := reportService.GetPublishedReports()

	json.NewEncoder(w).Encode(reportService.MapToS3Urls(reports))
}

func GetAllReports(w http.ResponseWriter, r *http.Request) {
	reports := models.GetAllReports()

	json.NewEncoder(w).Encode(reportService.MapToS3Urls(reports))
}
