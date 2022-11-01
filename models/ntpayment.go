package models

import (
	"api-ngmi/dbConn"
	"fmt"

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
	query := dbConn.DB.Model(payment).Where("access_wallet = ?", payment.AccessWallet)
	fmt.Println("This is the query as string", QueryToString(query))

	err := query.First()

	fmt.Println("This is the find error", err, payment.AccessWallet)

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
