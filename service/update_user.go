package service

import (
	"SmartHouseAPI/models"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (s *Service) UpdateUser(ctx context.Context, user models.User) error {
	_user, err := s.repo.ReadUser(ctx, user.Username)
	if err != nil {
		log.Println(err)
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(_user.Password), []byte(user.Password)) != nil {
		return fmt.Errorf("password is incorrect")
	}
	err = s.repo.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
