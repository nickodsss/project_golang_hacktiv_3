package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"assignment_3_golang/config"
	"assignment_3_golang/models"
)

func UpdateWeatherData() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		data := models.WeatherData{Water: water, Wind: wind}
		config.ConnectDatabase().Create(&data)

		waterStatus := waterStatus(water)
		windStatus := windStatus(wind)

		fmt.Printf("Water: %d meter, Status: %s | Wind: %d meter/sec, Status: %s\n", water, waterStatus, wind, windStatus)

		time.Sleep(15 * time.Second)
	}
}

func waterStatus(water int) string {
	if water < 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func windStatus(wind int) string {
	if wind < 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
