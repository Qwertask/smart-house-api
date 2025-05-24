package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetDoorSensorRecord(ctx context.Context, id int) (*models.DoorSensor, error) {
	record, err := s.repo.ReadDoorSensorRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
