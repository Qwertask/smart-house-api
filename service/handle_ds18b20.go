package service

import (
	"SmartHouseAPI/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Service) handleDs18b20(topic string, payload []byte) {
	var data models.DS18B20Sensor

	if err := json.Unmarshal(payload, &data); err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		return
	}
	data.Name = "ds18b20"

	log.Printf("📥 %s: %f @ %s\n", data.Name, data.Temperature, data.Timestamp.Format(time.RFC822Z))

	// Сохраняем в БД
	err := s.repo.CreateDs18b20SensorRecord(context.Background(), &models.DS18B20Sensor{
		Sensor: models.Sensor{
			Name: data.Name,
		},
		Temperature: data.Temperature,
		Timestamp:   data.Timestamp,
	})
	if err != nil {
		log.Println("❌ DB:", err)
	}
}
