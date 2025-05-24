package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) ReadWaterLeakSensorRecord(ctx context.Context, deviceID int) (*models.WaterLeakSensor, error) {
	query := `SELECT device_name, is_leaking, recorded_at FROM water_leak_readings WHERE device_id = $1 ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, deviceID)

	var record models.WaterLeakSensor
	if err := row.Scan(&record.ID, &record.IsLeaking, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
