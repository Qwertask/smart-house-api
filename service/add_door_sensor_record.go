package service

import (
	"SmartHouseAPI/models"
	"context"
	"fmt"
	"log"
	"time"
)

func (s *Service) AddDoorSensorRecord(ctx context.Context, record models.SetupDoorRequest) error {
	openDoor := true
	if record.Setup == "false" {
		openDoor = false
	}
	err := s.repo.CreateDoorSensorRecord(ctx, &models.DoorSensor{
		Sensor: models.Sensor{
			Name: fmt.Sprintf("door_%d", record.DoorID),
		},
		IsDoorOpen: openDoor,
		Timestamp:  time.Now(),
	})
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.SetDoor(ctx, openDoor, record.DoorID)
	if err != nil {
		//в случае провальной попытки отправить запрос на mqtt добавляем отмену
		err2 := s.repo.CreateDoorSensorRecord(ctx, &models.DoorSensor{
			Sensor: models.Sensor{
				Name: fmt.Sprintf("door_%d", record.DoorID),
			},
			IsDoorOpen: !openDoor,
			Timestamp:  time.Now(),
		})
		if err2 != nil {
			log.Println(err)
		}
		log.Println(err)
		return err
	}
	return nil
}
