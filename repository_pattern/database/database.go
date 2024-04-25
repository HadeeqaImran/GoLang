package database

import (
	"log"
	"os"
	"repo_pattern/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Filed to connect to database!\n ", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Add migrations
	db.AutoMigrate(&models.Task{})

	Database = DbInstance{Db: db}

}
