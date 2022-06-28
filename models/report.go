package models

import (
	"api-ngmi/dbConn"
	"log"
)

type ReportType int64

const (
	NFT ReportType = 0
)

type Chain int64

const (
	Ethereum Chain = 0
	BSC      Chain = 1
	Solana   Chain = 2
	Avax     Chain = 3
)

type SocialMedia struct {
	DiscordUrl string
	TwitterUrl string
	WebsiteUrl string
}

type MintDetails struct {
	Price        float64
	PresalePrice float64
	CoinTicker   string
}

type Report struct {
	Id                   int                    `pg:"id"`
	CreatedAt            string                 `pg:"created_at"`
	UpdatedAt            string                 `pg:"updated_at"`
	Name                 string                 `pg:"name"`
	Description          string                 `pg:"description"`
	Logo                 string                 `pg:"logo"`
	Chain                Chain                  `pg:"chain"`
	Socials              map[string]interface{} `pg:"socials"`
	ReportDetails        map[string]interface{} `pg:"report_details"`
	DetailedAnalysis     map[string]interface{} `pg:"detailed_analysis"`
	ReportDetailsLink    string                 `pg:"report_details_link"`
	DetailedAnalysisLink string                 `pg:"detailed_analysis_link"`
	Published            bool                   `pg:"published"`
}

func (report *Report) Save() {
	_, err := dbConn.DB.Model(report).Insert()
	if err != nil {
		log.Println("Can't insert", err)
	}
}

func (report *Report) Update() {
	_, err := dbConn.DB.Model(report).WherePK().Update()
	if err != nil {
		log.Println("Can't Update", err)
	}
}

func (report *Report) Find() {
	err := dbConn.DB.Model(report).WherePK().First()
	if err != nil {
		log.Println("Can't insert", err)
	}
}

func GetAllReports() []Report {
	got := []Report{}

	err := dbConn.DB.Model(&got).Select()
	if err != nil {
		log.Println("Can't retrieve Reports", err)
	}

	return got
}

func GetPublishedReports() []Report {
	got := []Report{}

	err := dbConn.DB.Model(&got).Where("published = ?", true).Select()
	if err != nil {
		log.Println("Can't retrieve Reports", err)
	}

	return got
}

func DeleteReportById(id int) bool {
	_, err := dbConn.DB.Model(&Report{}).Where("Id = ?", id).Delete()

	if err != nil {
		log.Println("Can't delete Report", err)
		return false
	}

	return true
}

func ApproveReview(id int) bool {
	_, err := dbConn.DB.Model(&Report{}).Set("published = true").Where("Id = ?", id).Update()

	if err != nil {
		log.Println("Can't publish Report", err)
		return false
	}

	return true
}
