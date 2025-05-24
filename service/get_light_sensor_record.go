package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetLightSensorRecord(ctx context.Context, id int) (*models.LightSensor, error) {
	record, err := s.repo.ReadLightSensorRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
