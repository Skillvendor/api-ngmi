package report

import (
	"api-ngmi/models"
	reportService "api-ngmi/services/report"
	"api-ngmi/types"
	"encoding/json"
	"errors"
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
			case 6: // LogoUrl
				fReport.LogoUrl = v.String()
			case 7: // Assets
				fReport.Assets = v.Interface().([]string)
			case 8: // AssetsUrls
				fReport.AssetsUrls = v.Interface().([]string)
			case 9: // Scores
				fReport.Scores = v.Interface().([]map[string]interface{})
			case 10: // Tags
				fReport.Tags = v.Interface().([]string)
			case 11: // Chain
				fReport.Chain = v.String()
			case 12: // Socials
				fReport.Socials = v.Interface().([]map[string]interface{})
			case 13: // ReportDetails
				fReport.ReportDetails = v.Interface().([]map[string]interface{})
			case 14: // DetailedAnalysis
				fReport.DetailedAnalysis = v.Interface().([]map[string]interface{})
			case 15: // ReportDetailsLink
				fReport.ReportDetailsLink = v.String()
			case 16: // DetailedAnalysisLink
				fReport.DetailedAnalysisLink = v.String()
			case 17: // Published
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

func CreateReport(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	err := json.NewDecoder(r.Body).Decode(&newReport)

	if err != nil {
		fmt.Println(err)
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't decode report"),
		}
	}
	newReport.Published = false

	saved := newReport.Save()

	if !saved {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't save report"),
		}
	}

	err = json.NewEncoder(w).Encode(newReport)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode report"),
		}
	}

	return nil
}

func UpdateReport(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error getting id"),
		}
	}

	err = json.NewDecoder(r.Body).Decode(&newReport)
	newReport.Id = id

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error decoding report"),
		}
	}

	updated := newReport.Update()

	if !updated {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't update report"),
		}
	}

	err = json.NewEncoder(w).Encode(newReport)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding report"),
		}
	}

	return nil
}

func GetReport(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error getting id"),
		}
	}

	found := newReport.Find()

	if !found {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't find report"),
		}
	}

	err = json.NewEncoder(w).Encode(filterReport(reportService.TransformToS3Urls(newReport), "gold"))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding S3 url transformations"),
		}
	}

	return nil
}

func DeleteReport(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error getting id"),
		}
	}

	deleted := newReport.Delete()

	if !deleted {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't delete report"),
		}
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func ApproveReview(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error encoding report"),
		}
	}

	published := newReport.Publish()

	if !published {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't publish report"),
		}
	}

	json.NewEncoder(w).Encode(newReport)

	return nil
}

func RejectReview(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("error encoding report"),
		}
	}

	rejected := newReport.Reject()

	if !rejected {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't reject report"),
		}
	}

	json.NewEncoder(w).Encode(newReport)

	return nil
}

func GetPublishedReports(w http.ResponseWriter, r *http.Request) error {
	reports := reportService.GetPublishedReports()

	// gigel := r.URL.Query()["tags"]

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(types.StandardError{Message: "Can't parse params"})
	// 	return
	// }

	err := json.NewEncoder(w).Encode(MapFilterReport(reportService.MapToS3Urls(reports), "gold", filterReport))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode reports"),
		}
	}

	return nil
}

// only for admin
func GetAllReports(w http.ResponseWriter, r *http.Request) error {
	reports := models.GetAllReports()

	err := json.NewEncoder(w).Encode(reportService.MapToS3Urls(reports))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode report"),
		}
	}

	return nil
}

func GetReportAdmin(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error getting it"),
		}
	}

	found := newReport.FindAdmin()

	if !found {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't find report"),
		}
	}

	err = json.NewEncoder(w).Encode(reportService.TransformToS3Urls(newReport))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error Encoding S3 url transformations"),
		}
	}

	return nil
}
