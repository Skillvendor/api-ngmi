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
	Id                   int                    `pg:"id" visibility:"free bronze silver gold" json:"id,omitempty"`
	CreatedAt            string                 `pg:"created_at" visibility:"free bronze silver gold" json:"created_at,omitempty"`
	UpdatedAt            string                 `pg:"updated_at" visibility:"free bronze silver gold" json:"updated_at,omitempty"`
	Name                 string                 `pg:"name" visibility:"free bronze silver gold" json:"name,omitempty"`
	Description          string                 `pg:"description" visibility:"free bronze silver gold" json:"description,omitempty"`
	Logo                 string                 `pg:"logo" visibility:"free bronze silver gold" json:"logo,omitempty"`
	Chain                string                 `pg:"chain" visibility:"free bronze silver gold" json:"chain,omitempty"`
	Socials              map[string]interface{} `pg:"socials" visibility:"bronze silver gold" json:"socials,omitempty"`
	ReportDetails        map[string]interface{} `pg:"report_details" visibility:"silver gold" json:"report_details,omitempty"`
	DetailedAnalysis     map[string]interface{} `pg:"detailed_analysis" visibility:"gold" json:"detailed_analysis,omitempty"`
	ReportDetailsLink    string                 `pg:"report_details_link" visibility:"silver gold" json:"report_details_link,omitempty"`
	DetailedAnalysisLink string                 `pg:"detailed_analysis_link" visibility:"gold" json:"detailed_analysis_link,omitempty"`
	Published            bool                   `pg:"published" visibility:"free bronze silver gold" json:"published,omitempty"`
}

func (report *Report) Save() bool {
	_, err := dbConn.DB.Model(report).Insert()
	if err != nil {
		log.Println("Can't insert", err)
		return false
	}

	return true
}

func (report *Report) Update() bool {
	_, err := dbConn.DB.Model(report).WherePK().Update()
	if err != nil {
		log.Println("Can't Update", err)
		return false
	}

	return true
}

func (report *Report) Delete() bool {
	_, err := dbConn.DB.Model(report).WherePK().Delete()
	if err != nil {
		log.Println("Can't Delete", err)
		return false
	}

	return true
}

func (report *Report) Find() bool {
	err := dbConn.DB.Model(report).WherePK().First()
	if err != nil {
		log.Println("Can't insert", err)
		return false
	}

	return true
}

func (report *Report) Publish() bool {
	_, err := dbConn.DB.Model(report).Set("published = true").WherePK().Update()

	if err != nil {
		log.Println("Can't publish Report", err)
		return false
	}

	return true
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
