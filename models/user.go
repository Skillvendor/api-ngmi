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
	AccessLevel int    `pg:"access_level" json:"access_level"`
	Username    string `pg:"username" json:"username"`
	Password    string `pg:"password" json:"-"`
}

func (user *User) Save() bool {
	_, err := dbConn.DB.Model(user).Insert()
	if err != nil {
		log.Println("Can't insert user", err)
		return false
	}

	return true
}

func (user *User) Find() bool {
	err := dbConn.DB.Model(user).Where("address = ?", user.Address).First()

	return err == nil
}

func (user *User) Update() bool {
	_, err := dbConn.DB.Model(user).
		Column("nonce", "auth_token").
		WherePK().
		Update()

	return err == nil
}

func (user *User) FindByUsername() bool {
	err := dbConn.DB.Model(user).Where("username = ?", user.Username).First()

	return err == nil
}
