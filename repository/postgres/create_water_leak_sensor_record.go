package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateWaterLeakRecord(ctx context.Context, record *models.WaterLeakSensor) error {
	query := `INSERT INTO water_leak_readings(device_name, is_leaking, recorded_at) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, record.Name, record.IsLeaking, record.Timestamp)
	return err
}
