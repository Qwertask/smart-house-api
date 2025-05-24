package service

import (
	"SmartHouseAPI/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Service) handleLeak1(topic string, payload []byte) {
	var data models.WaterLeakSensor

	if err := json.Unmarshal(payload, &data); err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		return
	}
	data.Name = "leak_1"

	log.Printf("📥 %s is leaking: %t @ %s\n", data.Name, data.IsLeaking, data.Timestamp.Format(time.RFC822Z))

	// Сохраняем в БД
	err := s.repo.CreateWaterLeakRecord(context.Background(), &models.WaterLeakSensor{
		Sensor: models.Sensor{
			Name: data.Name,
		},
		IsLeaking: data.IsLeaking,
		Timestamp: data.Timestamp,
	})
	if err != nil {
		log.Println("❌ DB:", err)
	}
}
