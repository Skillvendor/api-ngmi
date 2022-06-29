package report

import (
	"api-ngmi/models"
	reportService "api-ngmi/services/report"
	"api-ngmi/types"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"net/http"

	"reflect"
)

func filterReport(r models.Report, role string) models.Report {
	var fReport models.Report
	ev := reflect.ValueOf(r)
	et := reflect.TypeOf(r)

	// Iterate through each field within the struct
	for i := 0; i < ev.NumField(); i++ {
		v := ev.Field(i)
		t := et.Field(i)
		roles := t.Tag.Get("visibility")

		if strings.Contains(roles, role) {
			switch i {
			case 0: // ID
				fReport.Id = int(v.Int())
			case 1: // CreatedAt
				fReport.CreatedAt = v.String()
			case 2: // UpdatedAt
				fReport.UpdatedAt = v.String()
			case 3: // Name
				fReport.Name = v.String()
			case 4: // Description
				fReport.Description = v.String()
			case 5: // Logo
				fReport.Logo = v.String()
			case 6: // Chain
				fReport.Chain = v.String()
			case 7: // Socials
				fReport.Socials = v.Interface().(map[string]interface{})
			case 8: // ReportDetails
				fReport.ReportDetails = v.Interface().(map[string]interface{})
			case 9: // DetailedAnalysis
				fReport.DetailedAnalysis = v.Interface().(map[string]interface{})
			case 10: // ReportDetailsLink
				fReport.ReportDetailsLink = v.String()
			case 11: // DetailedAnalysisLink
				fReport.DetailedAnalysisLink = v.String()
			case 12: // Published
				fReport.Published = v.Bool()
			}
		}
	}

	return fReport
}

func MapFilterReport(vs []models.Report, role string, f func(models.Report, string) models.Report) []models.Report {
	vsm := make([]models.Report, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, role)
	}
	return vsm
}

func CreateReport(w http.ResponseWriter, r *http.Request) {
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
}

func UpdateReport(w http.ResponseWriter, r *http.Request) {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Getting id"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newReport)
	newReport.Id = id

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
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Getting Id"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newReport.Find()

	err = json.NewEncoder(w).Encode(filterReport(reportService.TransformToS3Urls(newReport), "free"))

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding S3 url transformations"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteReport(w http.ResponseWriter, r *http.Request) {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Decoding report"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newReport.Delete()

	json.NewEncoder(w).Encode([]byte("Ok!"))
}

func ApproveReview(w http.ResponseWriter, r *http.Request) {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		json.NewEncoder(w).Encode(types.StandardError{Message: "Error Encoding Token"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Do I find it?", newReport.Find())
	fmt.Println("IS IT PUBLISHED?", newReport.Publish())

	json.NewEncoder(w).Encode(newReport)
}

func GetPublishedReports(w http.ResponseWriter, r *http.Request) {
	reports := reportService.GetPublishedReports()

	json.NewEncoder(w).Encode(MapFilterReport(reportService.MapToS3Urls(reports), "free", filterReport))
}

// only for admin
func GetAllReports(w http.ResponseWriter, r *http.Request) {
	reports := models.GetAllReports()

	json.NewEncoder(w).Encode(reportService.MapToS3Urls(reports))
}
