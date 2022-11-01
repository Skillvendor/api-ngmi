package report

import (
	"api-ngmi/models"
	reportService "api-ngmi/services/report"
	"api-ngmi/types"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	reports := models.GetAllReports()

	err := json.NewEncoder(w).Encode(reportService.ProcessReports(reports))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode report"),
		}
	}

	return nil
}

func Show(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Admin Show")
	newReport := models.Report{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	newReport.Id = id

	if err != nil {
		fmt.Println("err while getting id", err)
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error getting it"),
		}
	}

	found := newReport.FindAdmin()

	if !found {
		fmt.Println("Report was not found")
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't find report"),
		}
	}

	fmt.Println("Got the report", newReport)

	err = json.NewEncoder(w).Encode(reportService.ProcessReport(newReport))

	if err != nil {
		fmt.Println("Error encoding", err)
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error Encoding S3 url transformations"),
		}
	}

	return nil
}

func Create(w http.ResponseWriter, r *http.Request) error {
	newReport := models.Report{}

	err := json.NewDecoder(r.Body).Decode(&newReport)

	if err != nil {
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

func Update(w http.ResponseWriter, r *http.Request) error {
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

func Delete(w http.ResponseWriter, r *http.Request) error {
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

func Publish(w http.ResponseWriter, r *http.Request) error {
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

func Unpublish(w http.ResponseWriter, r *http.Request) error {
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
