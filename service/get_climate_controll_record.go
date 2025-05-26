package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetClimateControlRecord(ctx context.Context, deviceName string) (*models.ClimateControl, error) {
	record, err := s.repo.ReadClimateControlRecord(ctx, deviceName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
