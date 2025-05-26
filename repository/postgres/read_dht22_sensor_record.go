package postgres

import (
	"SmartHouseAPI/models"
	"context"
	"fmt"
)

func (r *Repository) ReadDHT22SensorRecord(ctx context.Context, deviceID int) (*models.DHT22Sensor, error) {
	query := `SELECT device_name, temperature, humidity, recorded_at FROM dht22_readings WHERE device_name = $1 ORDER BY recorded_at DESC LIMIT 1`
	row := r.DB.QueryRowContext(ctx, query, fmt.Sprintf("dht22_%d", deviceID))

	var record models.DHT22Sensor
	if err := row.Scan(&record.ID, &record.Temperature, &record.Humidity, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return &record, nil
}
