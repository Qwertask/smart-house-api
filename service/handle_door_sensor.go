package service

import (
	"SmartHouseAPI/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Service) handleDoorSensor(topic string, payload []byte) {
	var data models.DoorSensor

	if err := json.Unmarshal(payload, &data); err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		return
	}
	data.Name = "door"

	log.Printf("📥 %s open: %t @ %s\n", data.Name, data.IsDoorOpen, data.Timestamp.Format(time.RFC822Z))

	// Сохраняем в БД
	err := s.repo.CreateDoorSensorRecord(context.Background(), &models.DoorSensor{
		Sensor: models.Sensor{
			Name: data.Name,
		},
		IsDoorOpen: data.IsDoorOpen,
		Timestamp:  data.Timestamp,
	})
	if err != nil {
		log.Println("❌ DB:", err)
	}
}
