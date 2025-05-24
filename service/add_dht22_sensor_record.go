package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddDHT22SensorRecord(ctx context.Context, record models.DHT22Sensor) error {
	err := s.repo.CreateDHT22SensorRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
