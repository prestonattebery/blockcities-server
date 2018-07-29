package db

import (
	"log"

	"github.com/devinroche/blockcities-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver for golang
)

//DB init gorm database
var DB *gorm.DB

// Open connects to postgres database
func Open() error {
	var err error
	DB, err = gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=blockcities sslmode=disable password=password")

	if err != nil {
		log.Fatal(err)
		return err
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Building{})
	return err
}

// Close ends database connection
func Close() error {
	return DB.Close()
}
