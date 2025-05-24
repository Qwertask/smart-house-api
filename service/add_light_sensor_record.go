package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddLightSensorRecord(ctx context.Context, record models.LightSensor) error {
	err := s.repo.CreateLightSensorRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
