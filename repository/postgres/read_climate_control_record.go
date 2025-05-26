package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) ReadClimateControlRecord(ctx context.Context, deviceName string) (*models.ClimateControl, error) {
	query := `SELECT device_name, is_on, target_temp, recorded_at FROM climate_control_states WHERE device_name = $1 ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, deviceName)

	var record models.ClimateControl
	if err := row.Scan(&record.ID, &record.IsOn, &record.TargetTemp, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
