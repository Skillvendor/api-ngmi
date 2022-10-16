package models

import (
	"api-ngmi/dbConn"
	"log"
)

const (
	Pending int = 1
	Success int = 2
	Fail    int = 3
)

type Transaction struct {
	Id            int    `pg:"id" json:"-"`
	Address       string `pg:"address" json:"address"`
	TransactionId string `pg:"transaction_id" json:"transaction_id"`
	State         int    `pg:"state" json:"state"`
}

func (transaction *Transaction) Save() bool {
	_, err := dbConn.DB.Model(transaction).Insert()
	if err != nil {
		log.Println("Can't insert transaction", err)
		return false
	}

	return true
}

func (transaction *Transaction) FindPending() bool {
	err := dbConn.DB.Model(transaction).Where("address = ?", transaction.Address).Where("state = ?", Pending).First()

	return err == nil
}

func (transaction *Transaction) Update() bool {
	_, err := dbConn.DB.Model(transaction).
		Column("state").
		WherePK().
		Update()

	return err == nil
}

func (transaction *Transaction) Find() bool {
	err := dbConn.DB.Model(transaction).WherePK().First()

	return err == nil
}
