package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) ReadLightSensorRecord(ctx context.Context, deviceID int) (*models.LightSensor, error) {
	query := `SELECT device_name, is_on, brightness, updated_at FROM light_readings WHERE device_id = $1 ORDER BY updated_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, deviceID)

	var record models.LightSensor
	if err := row.Scan(&record.ID, &record.IsOn, &record.Brightness, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
