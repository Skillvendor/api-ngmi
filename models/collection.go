package models

import (
	"go-rarity/dbConn"
	"log"
)

type ReportType int64

const (
	NFT ReportType = 0
)

type Tier int64

const (
	Free Tier = 0
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

type Collection struct {
	Id          int         `pg:"id"`
	Name        string      `pg:"name"`
	Description string      `pg:"description"`
	Socials     SocialMedia `pg:"socials"`
	MintTime    string      `pg:"mint_time"`
	Logos       []string    `pg:"logos"`
	Research    []string    `pg:"research"`
	MintInfo    MintDetails `pg:"mint_info"`
	ReportType  ReportType  `pg:"report_type"`
	Tier        Tier        `pg:"tier"`
	Published   bool        `pg:"published"`
}

var TableName = "collections"

func (collection *Collection) Save() {
	_, err := dbConn.DB.Model(collection).Insert()
	if err != nil {
		log.Println("Can't insert", err)
	}
}

func (collection *Collection) Update() {
	_, err := dbConn.DB.Model(collection).WherePK().Update()
	if err != nil {
		log.Println("Can't Update", err)
	}
}

func (collection *Collection) Find() {
	err := dbConn.DB.Model(collection).WherePK().First()
	if err != nil {
		log.Println("Can't insert", err)
	}
}

func GetAllCollections() []Collection {
	got := []Collection{}

	err := dbConn.DB.Model(&got).Select()
	if err != nil {
		log.Println("Can't retrieve collections", err)
	}

	return got
}

func GetPublishedCollections() []Collection {
	got := []Collection{}

	err := dbConn.DB.Model(&got).Where("published = ?", true).Select()
	if err != nil {
		log.Println("Can't retrieve collections", err)
	}

	return got
}

func DeleteCollectionById(id int) bool {
	_, err := dbConn.DB.Model(&Collection{}).Where("Id = ?", id).Delete()

	if err != nil {
		log.Println("Can't delete collection", err)
		return false
	}

	return true
}

func ApproveReview(id int) bool {
	_, err := dbConn.DB.Model(&Collection{}).Set("published = true").Where("Id = ?", id).Update()

	if err != nil {
		log.Println("Can't publish collection", err)
		return false
	}

	return true
}
