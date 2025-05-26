package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetDS18B20SensorRecord(ctx context.Context) (*models.DS18B20Sensor, error) {
	record, err := s.repo.ReadDs18b20SensorRecord(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
