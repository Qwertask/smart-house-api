package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateClimateControlRecord(ctx context.Context, record *models.ClimateControl) error {
	query := `INSERT INTO climate_control_states(device_name, is_on, target_temp) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, record.Name, record.IsOn, record.TargetTemp)
	return err
}
