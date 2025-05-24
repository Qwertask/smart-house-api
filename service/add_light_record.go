package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) AddLightRecord(ctx context.Context, record models.Light) error {
	err := s.repo.CreateLightRecord(ctx, &record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
