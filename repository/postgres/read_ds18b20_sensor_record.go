package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

// fixme:device name hardcoded
func (r *Repository) ReadDs18b20SensorRecord(ctx context.Context) (*models.DS18B20Sensor, error) {
	query := `SELECT device_name, temperature, recorded_at FROM ds18b20_readings WHERE device_name = 'ds18b20' ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query)

	var record models.DS18B20Sensor
	if err := row.Scan(&record.Name, &record.Temperature, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
