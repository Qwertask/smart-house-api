package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetDHT22SensorRecord(ctx context.Context, id int) (*models.DHT22Sensor, error) {
	record, err := s.repo.ReadDHT22SensorRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
