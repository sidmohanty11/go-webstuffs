package db

import (
	"log"

	"github.com/sidmohanty11/go-webstuffs/setting-up-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Sidharth11 dbname=testDb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	DB = db

	db.AutoMigrate(&models.User{})
}
