package models

import "time"

type WaterLeakSensor struct {
	Sensor
	IsLeaking bool      `json:"is_leaking"`
	Timestamp time.Time `json:"timestamp"`
}
