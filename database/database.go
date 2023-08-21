package database

import (
	"log"
	"os"

	"github.com/leonfiasco/go-auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DbInstance struct {
	Db *gorm.DB
}


var Database DbInstance


func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("goauth.db"), &gorm.Config{})


	if err != nil {
		log.Fatal("failed to connect to the database! \n", err)
		os.Exit(2)
	}


	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	
	// // TODO: Add migrations
	 db.AutoMigrate(&models.User{})

	Database = DbInstance{Db: db}
}