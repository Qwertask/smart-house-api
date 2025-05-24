package service

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

func (s *Service) GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(s.jwt.Expire).Unix()
	tokenString, err := token.SignedString([]byte(s.jwt.Key))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
