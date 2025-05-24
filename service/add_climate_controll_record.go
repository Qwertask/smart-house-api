package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddClimateControlRecord(ctx context.Context, record models.ClimateControl) error {
	err := s.repo.CreateClimateControlRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
