package service

import (
	"errors"
	"log"
	"time"
)

func (s *Service) ValidateJWTToken(tokenString string) (bool, error) {
	_, claims, err := s.parseJWTToken(tokenString)
	if err != nil {
		log.Println(err)
		return false, err
	}

	expiration, ok := claims["exp"].(float64)
	if !ok {
		return false, errors.New("exp field is not float64")
	}
	if time.Now().Unix() > int64(expiration) {
		return false, errors.New("token is expired")
	}
	return true, nil
}
