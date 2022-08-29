package models

import (
	"api-ngmi/dbConn"
	"log"
)

type Admin struct {
	Id          int    `pg:"id" json:"-"`
	Username    string `pg:"username" json:"username,omitempty"`
	Password    string `pg:"password" json:"password,omitempty"`
	AuthToken   string `pg:"auth_token" json:"-"`
	AccessLevel int    `pg:"access_level" json:"-"`
}

func (admin *Admin) Save() bool {
	_, err := dbConn.DB.Model(admin).Insert()
	if err != nil {
		log.Println("Can't insert admin", err)
		return false
	}

	return true
}

func (admin *Admin) Find() bool {
	err := dbConn.DB.Model(admin).Where("username = ?", admin.Username).First()

	return err == nil
}

func (admin *Admin) Update() bool {
	_, err := dbConn.DB.Model(admin).
		Column("auth_token").
		WherePK().
		Update()

	return err == nil
}
