package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateDs18b20SensorRecord(ctx context.Context, record *models.DS18B20Sensor) error {
	query := `INSERT INTO ds18b20_readings(device_name, temperature, recorded_at) VALUES ($1, $2, $4)`
	_, err := r.DB.ExecContext(ctx, query, record.Name, record.Temperature, record.Timestamp)
	return err
}
