package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetWaterLeakSensorRecord(ctx context.Context, id int) (*models.WaterLeakSensor, error) {
	record, err := s.repo.ReadWaterLeakSensorRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
