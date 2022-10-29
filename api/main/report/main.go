package report

import (
	"api-ngmi/constants"
	"api-ngmi/models"
	reportService "api-ngmi/services/report"
	userService "api-ngmi/services/user"
	"api-ngmi/types"
	"encoding/json"
	"errors"
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
				fReport.Scores = v.Interface().([]models.Score)
			case 10: // Average Score
				fReport.AverageScore = float64(v.Float())
			case 11: // Tags
				fReport.Tags = v.Interface().([]string)
			case 12: // Chain
				fReport.Chain = v.String()
			case 13: // Socials
				fReport.Socials = v.Interface().([]map[string]interface{})
			case 14: // ReportDetails
				fReport.ReportDetails = v.Interface().(map[string]interface{})
			case 15: // DetailedAnalysis
				fReport.DetailedAnalysis = v.Interface().([]map[string]interface{})
			case 16: // ReportDetailsLink
				fReport.ReportDetailsLink = v.String()
			case 17: // DetailedAnalysisLink
				fReport.DetailedAnalysisLink = v.String()
			case 18: // Published
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

func Index(w http.ResponseWriter, r *http.Request) error {
	reports := reportService.GetPublishedReports()

	var user models.User = r.Context().Value("user").(models.User)

	err := json.NewEncoder(w).Encode(MapFilterReport(reportService.ProcessReports(reports), constants.LevelToTier[userService.AccessLevelFor(user.Address)], filterReport))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode reports"),
		}
	}

	return nil
}

func Show(w http.ResponseWriter, r *http.Request) error {
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

	var user models.User = r.Context().Value("user").(models.User)

	err = json.NewEncoder(w).Encode(filterReport(reportService.ProcessReport(newReport), constants.LevelToTier[user.AccessLevel]))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding S3 url transformations"),
		}
	}

	return nil
}
