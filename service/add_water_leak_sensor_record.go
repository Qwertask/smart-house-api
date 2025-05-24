package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddWaterLeakSensorRecord(ctx context.Context, record models.WaterLeakSensor) error {
	err := s.repo.CreateWaterLeakRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
