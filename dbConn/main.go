package dbConn

import (
	"log"
	"os"

	pg "github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

var DB *pg.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connUrl := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	DB = pg.Connect(&pg.Options{
		Addr:     connUrl,
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})
}
