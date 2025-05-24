package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddDoorSensorRecord(ctx context.Context, record models.DoorSensor) error {
	err := s.repo.CreateDoorSensorRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
