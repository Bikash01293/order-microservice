package db

import (
	"log"
	"order-micro/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler{
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	db.AutoMigrate(&models.Order{})

	return Handler{db}
}