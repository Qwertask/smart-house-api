package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetMotionSensorRecord(ctx context.Context, id int) (*models.MotionSensor, error) {
	record, err := s.repo.ReadMotionSensorRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
