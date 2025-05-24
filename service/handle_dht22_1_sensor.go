package service

import (
	"SmartHouseAPI/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Service) handleDHT221(topic string, payload []byte) {
	var data models.DHT22Sensor

	if err := json.Unmarshal(payload, &data); err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		return
	}
	data.Name = "dht22_1"

	log.Printf("📥 %s: %f C	%f%% @ %s\n", data.Name, data.Temperature, data.Humidity, data.Timestamp.Format(time.RFC822Z))

	// Сохраняем в БД
	err := s.repo.CreateDHT22SensorRecord(context.Background(), &models.DHT22Sensor{
		Sensor: models.Sensor{
			Name: data.Name,
		},
		Temperature: data.Temperature,
		Humidity:    data.Humidity,
		Timestamp:   data.Timestamp,
	})
	if err != nil {
		log.Println("❌ DB:", err)
	}
}
