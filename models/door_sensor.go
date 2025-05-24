package models

import "time"

type DoorSensor struct {
	Sensor
	IsDoorOpen bool      `json:"is_door_open"` // для
	Timestamp  time.Time `json:"timestamp"`
}
