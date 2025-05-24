package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) ReadDs18b20SensorRecord(ctx context.Context, deviceID string) (*models.DS18B20Sensor, error) {
	query := `SELECT device_name, temperature, recorded_at FROM dht22_readings WHERE device_id = $1 ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, deviceID)

	var record models.DS18B20Sensor
	if err := row.Scan(&record.Name, &record.Temperature, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
