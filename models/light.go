package models

import "time"

type Light struct {
	Sensor
	IsOn       bool      `json:"is_on"`
	Brightness int       `json:"brightness"` // 0-100%
	Timestamp  time.Time `json:"timestamp"`
}
