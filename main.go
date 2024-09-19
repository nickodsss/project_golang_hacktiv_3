package main

import (
	"assignment_3_golang/config"
	"assignment_3_golang/controllers"
)

func main() {
	config.ConnectDatabase()

	controllers.UpdateWeatherData()
}
