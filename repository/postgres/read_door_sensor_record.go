package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) ReadDoorSensorRecord(ctx context.Context) (*models.DoorSensor, error) {
	query := `SELECT device_name, is_open, recorded_at FROM door_states WHERE device_name = 'door' ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query)

	var record models.DoorSensor
	if err := row.Scan(&record.ID, &record.IsDoorOpen, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
