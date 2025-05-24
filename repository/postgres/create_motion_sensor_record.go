package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateMotionRecord(ctx context.Context, record *models.MotionSensor) error {
	query := `INSERT INTO motion_readings(device_name, motion, recorded_at) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, record.Name, record.IsMotion, record.Timestamp)
	return err
}
