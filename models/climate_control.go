package models

type ClimateControl struct {
	Sensor
	IsOn       bool    `json:"is_on"`
	TargetTemp float64 `json:"target_temp"` // целевая температура
}
