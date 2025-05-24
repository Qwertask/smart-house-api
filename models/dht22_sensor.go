package models

import "time"

type DHT22Sensor struct {
	Sensor
	Temperature float64   `json:"temperature"` // в °C
	Humidity    float64   `json:"humidity"`    // в %
	Timestamp   time.Time `json:"timestamp"`
}
