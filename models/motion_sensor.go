package models

import "time"

type MotionSensor struct {
	Sensor
	IsMotion  bool      `json:"is_motion"` // для PIR
	Timestamp time.Time `json:"timestamp"`
}
