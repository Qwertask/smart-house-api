package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetClimateControlRecord(ctx context.Context, id int) (*models.ClimateControl, error) {
	record, err := s.repo.ReadClimateControlRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
