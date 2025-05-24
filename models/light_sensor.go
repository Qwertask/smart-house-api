package models

type LightSensor struct {
	Sensor
	IsOn       bool `json:"is_on"`
	Brightness int  `json:"brightness"` // 0-100%
}
