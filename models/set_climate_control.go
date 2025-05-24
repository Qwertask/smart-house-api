package models

type SetClimateControl struct {
	IsOn              bool    `json:"is_on"`
	TargetTemperature float64 `json:"target_temperature"`
}
