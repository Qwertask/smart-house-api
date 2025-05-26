package service

import (
	"SmartHouseAPI/models"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, user models.User) (string, string, error) {
	userFromDB, err := s.repo.ReadUser(ctx, user.Username)
	if err != nil {
		return "", "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	if string(hashedPassword) != userFromDB.Password {
		return "", "", errors.New("invalid password")
	}
	accessToken, err := s.GenerateJWTAccessToken()
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.GenerateJWTRefreshToken()
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
