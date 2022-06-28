package models

import (
	"api-ngmi/dbConn"
	"log"
)

type User struct {
	Id          int    `pg:"id" json:"-"`
	Address     string `pg:"address" json:"address,omitempty"`
	Nonce       string `pg:"nonce" json:"nonce,omitempty"`
	AuthToken   string `pg:"auth_token" json:"-"`
	AccessLevel int    `pg:"access_level" json:"-"`
}

func (user *User) Save() {
	_, err := dbConn.DB.Model(user).Insert()
	if err != nil {
		log.Println("Can't insert user", err)
	}
}

func (user *User) Find() {
	err := dbConn.DB.Model(user).Where("address = ?", user.Address).First()
	if err != nil {
		log.Println("Can't find", err)
	}
}

func (user *User) Update() {
	_, err := dbConn.DB.Model(user).
		Column("Nonce", "AuthToken").
		WherePK().
		Update()
	if err != nil {
		log.Println("Can't Update", err)
	}
}
