package models

type WeatherData struct {
	ID    uint `gorm:"primaryKey"`
	Water int
	Wind  int
}
