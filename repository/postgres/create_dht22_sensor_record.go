package postgres

import (
	"SmartHouseAPI/models"
	"context"
)

func (r *Repository) CreateDHT22SensorRecord(ctx context.Context, record *models.DHT22Sensor) error {
	query := `INSERT INTO dht22_readings(device_name, temperature, humidity, recorded_at) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.ExecContext(ctx, query, record.Name, record.Temperature, record.Humidity, record.Timestamp)
	return err
}
