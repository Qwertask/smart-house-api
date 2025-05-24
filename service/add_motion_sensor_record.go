package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddMotionSensorRecord(ctx context.Context, record models.MotionSensor) error {
	err := s.repo.CreateMotionRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
