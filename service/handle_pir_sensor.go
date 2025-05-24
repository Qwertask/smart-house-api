package service

import (
	"SmartHouseAPI/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Service) handlePIR(topic string, payload []byte) {
	var data models.MotionSensor

	if err := json.Unmarshal(payload, &data); err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		return
	}
	data.Name = "pir"

	log.Printf("📥 %s motion detected: %t @ %s\n", data.Name, data.IsMotion, data.Timestamp.Format(time.RFC822Z))

	// Сохраняем в БД
	err := s.repo.CreateMotionRecord(context.Background(), &models.MotionSensor{
		Sensor: models.Sensor{
			Name: data.Name,
		},
		IsMotion:  data.IsMotion,
		Timestamp: data.Timestamp,
	})
	if err != nil {
		log.Println("❌ DB:", err)
	}
}
