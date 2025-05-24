package service

import (
	"SmartHouseAPI/models"
	"context"
	"log"
)

// GetUser
// Возвращает объект models.User, в котором в поле password находится хэш пароля
func (s *Service) GetUser(ctx context.Context, username string) (*models.User, error) {
	record, err := s.repo.ReadUser(ctx, username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return record, nil
}
