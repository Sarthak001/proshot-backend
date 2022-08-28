package db

import (
	"log"

	"github.com/Sarthak001/proshot-backend/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.UserDetails{}, &models.Album{}, &models.Photo{}, &models.SharedAlbum{})

	return db
}
