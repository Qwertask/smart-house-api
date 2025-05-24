package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

func (s *Service) GetLightRecord(ctx context.Context, id int) (*models.Light, error) {
	record, err := s.repo.ReadLightRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
