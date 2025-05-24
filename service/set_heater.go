package service

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Service) SetHeater(ctx context.Context, on bool, temperature float64) error {
	if temperature < 5 || temperature > 35 {
		return fmt.Errorf("heater temperature must be between 5 and 35 °C")
	}

	topic := "sensor/heater/command"

	cmd := struct {
		State       string  `json:"state"`
		Temperature float64 `json:"temperature"`
	}{
		State:       "off",
		Temperature: 0,
	}
	if on {
		cmd.State = "on"
		cmd.Temperature = temperature
	}

	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	return s.MQTT.Publish(topic, 1, false, payload)
}
