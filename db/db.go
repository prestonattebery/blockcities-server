package db

import (
	"fmt"
	"log"
	"os"

	"github.com/devinroche/blockcities-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver for golang
	"github.com/joho/godotenv"
)

//DB init gorm database
var DB *gorm.DB

// Open connects to postgres database
func Open() error {
	var err error

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	DB, err = gorm.Open("postgres", dbUri)

	if err != nil {
		log.Fatal(err)
		return err
	}

	DB.AutoMigrate(&models.User{}, &models.Building{})
	return err
}

// Close ends database connection
func Close() error {
	return DB.Close()
}
