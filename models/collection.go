package models

import (
	"go-rarity/dbConn"
	"log"
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
	Name        string      `pg:"name"`
	Description string      `pg:"description"`
	Socials     SocialMedia `pg:"socials"`
	MintTime    string      `pg:"mint_time"`
	Logos       []string    `pg:"logos"`
	Research    []string    `pg:"research"`
	MintInfo    MintDetails `pg:"mint_info"`
}

var TableName = "collections"

func (collection *Collection) Save() {
	_, err := dbConn.DB.Model(collection).Insert()
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
