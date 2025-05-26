package models

type SetupDoorRequest struct {
	DoorID int    `json:"door_id"`
	Setup  string `json:"setup"`
}
