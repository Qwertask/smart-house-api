package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateLightRecord(ctx context.Context, record *models.Light) error {
	query := `INSERT INTO light_readings(device_name, is_on, brightness) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, record.ID, record.IsOn, record.Brightness)
	return err
}
