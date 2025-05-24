package service

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Service) SetLight(ctx context.Context, lightID int, on bool, brightness int) error {
	if lightID < 1 || lightID > 12 {
		return fmt.Errorf("invalid lightID %d", lightID)
	}
	if brightness < 0 || brightness > 100 {
		return fmt.Errorf("brightness must be between 0 and 100")
	}

	topic := fmt.Sprintf("sensor/light_%d/command", lightID)

	cmd := struct {
		State      string `json:"state"`
		Brightness int    `json:"brightness"`
	}{
		State:      "off",
		Brightness: 0,
	}
	if on {
		cmd.State = "on"
		cmd.Brightness = brightness
	}

	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	return s.MQTT.Publish(topic, 1, false, payload)
}
