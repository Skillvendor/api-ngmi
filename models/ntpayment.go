package models

import (
	"api-ngmi/dbConn"
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10/orm"
)

type NtPayment struct {
	Id              int    `pg:"id" json:"-"`
	CitizenWallet   string `pg:"citizen_wallet" json:"citizen_wallet"`
	AccessWallet    string `pg:"access_wallet" json:"access_wallet"`
	TransactionLink string `pg:"transaction_link" json:"transaction_link"`

	DiscordHandle string `pg:"discord_handle" json:"discord_handle"`
	TwitterHandle string `pg:"twitter_handle" json:"twitter_handle"`

	TransactionVerified bool `pg:"transaction_verified" json:"transaction_verified"`
	TransactionValid    bool `pg:"transaction_valid" json:"transaction_valid"`
	Valid               bool `pg:"valid" json:"valid"`
}

func QueryToString(q *orm.Query) string {
	value, _ := q.AppendQuery(orm.NewFormatter(), nil)

	return string(value)
}

func (payment *NtPayment) Find() bool {
	err := dbConn.DB.Model(payment).Where("lower(access_wallet) = ?", strings.ToLower(payment.AccessWallet)).First()

	return err == nil
}

func GetAllPayments() []NtPayment {
	got := []NtPayment{}

	err := dbConn.DB.Model(&got).Select()
	if err != nil {
		fmt.Println("Can't retrieve Payments", err)
	}

	return got
}

func GetAllPendingValidationPayments() []NtPayment {
	got := []NtPayment{}

	query := dbConn.DB.Model(&got).Where("transaction_verified = ?", false)

	err := query.Select()
	if err != nil {
		fmt.Println("Can't retrieve Payments", err)
	}

	return got
}

func GetAllValidPayments() []NtPayment {
	got := []NtPayment{}

	query := dbConn.DB.Model(&got).Where("valid = ?", true)

	err := query.Select()
	if err != nil {
		fmt.Println("Can't retrieve Payments", err)
	}

	return got
}
