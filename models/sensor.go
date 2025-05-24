package models

import "time"

type Sensor struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"` // "floor1", "floor2"
	Type      string    `json:"type"`     // "temperature", "leak" и т.д.
	UpdatedAt time.Time `json:"updated_at"`
}
