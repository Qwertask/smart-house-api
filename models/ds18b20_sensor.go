package models

import "time"

type DS18B20Sensor struct {
	Sensor
	Temperature float64   `json:"temperature"` // в °C
	Timestamp   time.Time `json:"timestamp"`
}
