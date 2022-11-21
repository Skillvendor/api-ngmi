package main

import (
	"api-ngmi/models"
	"api-ngmi/services/auth"
)

func main() {
	adminUsername := "Skillvendor"
	adminPassword := "123456`"

	hashedPass, _ := auth.HashPassword(adminPassword)

	admin := models.Admin{
		Username: adminUsername,
		Password: hashedPass,
	}

	admin.Save()
}
