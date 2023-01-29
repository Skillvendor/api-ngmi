package models

import (
	"api-ngmi/dbConn"
	"fmt"
	"log"
	"strings"
	"time"
)

type CommunityPayment struct {
	Id              int    `pg:"id" json:"-"`
	CitizenWallet   string `pg:"citizen_wallet" json:"citizen_wallet"`
	AccessWallet    string `pg:"access_wallet" json:"access_wallet"`
	TransactionLink string `pg:"transaction_link" json:"transaction_link"`
	DiscordHandle   string `pg:"discord_handle" json:"discord_handle"`

	TransactionVerified bool `pg:"transaction_verified" json:"transaction_verified"`
	TransactionValid    bool `pg:"transaction_valid" json:"transaction_valid"`
	Valid               bool `pg:"valid" json:"valid"`

	ExpiresAt     time.Time `pg:"expires_at" json:"expires_at"`
	CommunityName string    `pg:"community_name" json:"community_name"`
	AccessLevel   string    `pg:"access_level" json:"access_level"`
}

func (payment *CommunityPayment) Save() bool {
	_, err := dbConn.DB.Model(payment).Insert()
	if err != nil {
		log.Println("Can't insert CommunityPayment", err)
		return false
	}

	return true
}

func (payment *CommunityPayment) Update() bool {
	_, err := dbConn.DB.Model(payment).Column("transaction_verified", "transaction_valid", "valid").WherePK().Update()
	if err != nil {
		log.Println("Can't insert CommunityPayment", err)
		return false
	}

	return true
}

func (payment *CommunityPayment) Find() bool {
	err := dbConn.DB.Model(payment).Where("lower(access_wallet) = ?", strings.ToLower(payment.AccessWallet)).First()

	fmt.Println("Are there errors?", err)
	return err == nil
}

func GetAllPaymentsCommunity() []CommunityPayment {
	got := []CommunityPayment{}

	err := dbConn.DB.Model(&got).Select()
	if err != nil {
		log.Println("Can't retrieve Payments", err)
	}

	return got
}

func GetAllPendingValidationPaymentsCommunity() []CommunityPayment {
	got := []CommunityPayment{}

	query := dbConn.DB.Model(&got).Where("transaction_verified = ?", false)

	err := query.Select()
	if err != nil {
		log.Println("Can't retrieve Payments", err)
	}

	return got
}

// func (payment *NtPayment) Update() bool {
// 	_, err := dbConn.DB.Model(payment).
// 		Column("nonce", "auth_token").
// 		WherePK().
// 		Update()

// 	return err == nil
// }
