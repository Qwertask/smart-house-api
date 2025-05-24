package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateDoorSensorRecord(ctx context.Context, record *models.DoorSensor) error {
	query := `INSERT INTO door_states(device_name, is_open, recorded_at) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, record.Name, record.IsDoorOpen, record.Timestamp)
	return err
}
