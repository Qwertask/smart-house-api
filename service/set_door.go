package service

import (
	"context"
	"encoding/json"
	"fmt"
)

// true - открыть дверь
// false - закрыть дверь
// булевы значения для оптиизации
func (s *Service) SetDoor(ctx context.Context, setup bool, doorId int) error {
	topic := fmt.Sprintf("door_%d/command", doorId)

	cmd := struct {
		State bool `json:"state"`
	}{
		State: setup,
	}
	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	return s.MQTT.Publish(topic, 1, false, payload)
}
