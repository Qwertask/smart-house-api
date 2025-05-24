package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) ReadMotionSensorRecord(ctx context.Context, deviceID int) (*models.MotionSensor, error) {
	query := `SELECT device_name, motion, recorded_at FROM motion_readings WHERE device_id = $1 ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, deviceID)

	var record models.MotionSensor
	if err := row.Scan(&record.ID, &record.IsMotion, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
