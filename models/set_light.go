package models

type SetLight struct {
	ID         int  `json:"id"`
	IsOn       bool `json:"is_on"`
	Brightness int  `json:"brightness"`
}
