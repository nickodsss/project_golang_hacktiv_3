package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"assignment_3_golang/models"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=assignment_3_golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.WeatherData{})

	return database
}
