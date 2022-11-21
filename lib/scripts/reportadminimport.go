package main

import (
	"api-ngmi/models"
	"fmt"
	"os"
	"strconv"

	"encoding/csv"
	"encoding/json"
)

func createReport(data [][]string) []models.Report {
	var reports []models.Report

	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Report
			for j, field := range line {
				if j == 0 {
					fmt.Println("Starting Report with id", field)
				} else if j == 1 {
					rec.Name = field
				} else if j == 2 {
					rec.Description = field
				} else if j == 3 {
					rec.Logo = field
				} else if j == 4 {
					rec.Assets = nil
					if len(field) > 0 {
						var dat []string
						if err := json.Unmarshal([]byte(field), &dat); err != nil {
							panic(err)
						}
						fmt.Println(dat)
						rec.Assets = dat
					}
				} else if j == 5 {
					rec.Chain = field
				} else if j == 6 {
					fmt.Println("Rating is not used")
				} else if j == 7 {
					fmt.Println("Rating breakdown is not used")
				} else if j == 8 {
					rec.Tags = nil
					if len(field) > 0 {
						var dat []string
						if err := json.Unmarshal([]byte(field), &dat); err != nil {
							panic(err)
						}
						rec.Tags = dat
					}
				} else if j == 9 {
					rec.Scores = nil
					if len(field) > 0 {
						var dat []models.Score
						if err := json.Unmarshal([]byte(field), &dat); err != nil {
							panic(err)
						}
						rec.Scores = dat
					}
				} else if j == 10 {
					rec.Socials = nil
					if len(field) > 0 {
						var dat []map[string]interface{}
						if err := json.Unmarshal([]byte(field), &dat); err != nil {
							panic(err)
						}
						rec.Socials = dat
					}
				} else if j == 11 {
					rec.ReportDetails = nil
					if len(field) > 0 {
						var dat map[string]interface{}
						if err := json.Unmarshal([]byte(field), &dat); err != nil {
							panic(err)
						}
						rec.ReportDetails = dat
					}
				} else if j == 12 {
					rec.DetailedAnalysis = nil
					if len(field) > 0 {
						var dat []map[string]interface{}
						if err := json.Unmarshal([]byte(field), &dat); err != nil {
							panic(err)
						}
						fmt.Println(dat)
						rec.DetailedAnalysis = dat
					}
				} else if j == 13 {
					rec.ReportDetailsLink = field
				} else if j == 14 {
					rec.DetailedAnalysisLink = field
				} else if j == 15 {
					rec.CreatedAt = field
				} else if j == 16 {
					rec.UpdatedAt = field
				} else if j == 17 {
					ok, err := strconv.ParseBool(field)
					if err != nil {
						fmt.Println("Can't Parse Bool")
					}
					rec.Published = ok
				}
			}
			rec.Save()
			reports = append(reports, rec)
		}
	}
	return reports
}

func createAdmin(data [][]string) []models.Admin {
	var admins []models.Admin

	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Admin
			for j, field := range line {
				if j == 0 {
					fmt.Println("Starting admin with id", field)
				} else if j == 1 {
					fmt.Println("Field 1 is", field)
					rec.Username = field
				} else if j == 2 {
					rec.Password = field
				} else if j == 3 {
					rec.AuthToken = field
				} else if j == 4 {
					fmt.Println("created at", field)
				} else if j == 5 {
					fmt.Println("updated at", field)
				} else if j == 6 {
					number, err := strconv.Atoi(field)
					if err != nil {
						fmt.Println("Can't parse OMG")
					}
					rec.AccessLevel = number
				}
			}
			rec.Save()
			admins = append(admins, rec)
		}
	}
	return admins
}

func Import() {
	// httpServer.Run()

	// filename := "admin_rows.csv"
	filename := "reports_rows.csv"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(createAdmin(data))
	fmt.Println(createReport(data))
}
