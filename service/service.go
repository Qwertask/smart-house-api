package service

import (
	"SmartHouseAPI/config"
	"SmartHouseAPI/internal/mqtt"
	"SmartHouseAPI/models"
	"context"
)

type Repository interface {
	CreateClimateControlRecord(ctx context.Context, record *models.ClimateControl) error
	CreateDHT22SensorRecord(ctx context.Context, record *models.DHT22Sensor) error
	CreateDoorSensorRecord(ctx context.Context, record *models.DoorSensor) error
	CreateLightRecord(ctx context.Context, record *models.Light) error
	CreateLightSensorRecord(ctx context.Context, record *models.LightSensor) error
	CreateMotionRecord(ctx context.Context, record *models.MotionSensor) error
	CreateWaterLeakRecord(ctx context.Context, record *models.WaterLeakSensor) error
	CreateDs18b20SensorRecord(ctx context.Context, record *models.DS18B20Sensor) error
	ReadClimateControlRecord(ctx context.Context, deviceID int) (*models.ClimateControl, error)
	ReadDHT22SensorRecord(ctx context.Context, deviceID int) (*models.DHT22Sensor, error)
	ReadDoorSensorRecord(ctx context.Context, deviceID int) (*models.DoorSensor, error)
	ReadLightRecord(ctx context.Context, deviceID int) (*models.Light, error)
	ReadLightSensorRecord(ctx context.Context, deviceID int) (*models.LightSensor, error)
	ReadMotionSensorRecord(ctx context.Context, deviceID int) (*models.MotionSensor, error)
	ReadUser(ctx context.Context, username string) (*models.User, error)
	ReadWaterLeakSensorRecord(ctx context.Context, deviceID int) (*models.WaterLeakSensor, error)
	UpdateUser(ctx context.Context, user models.User) error
}

type Service struct {
	MQTT *mqtt.MQTTService
	repo Repository
	jwt  config.JWTConfig
}

func NewService(mqttBroker string, r Repository, mqttConfig config.MQTTConfig) *Service {
	s := &Service{
		repo: r,
	}

	s.MQTT = mqtt.New(mqttBroker, mqttConfig.Password, mqttConfig.Login)
	s.MQTT.RegisterHandler("sensor/dht22_1/data", s.handleDHT221)
	s.MQTT.RegisterHandler("sensor/dht22_2/data", s.handleDHT221)
	s.MQTT.RegisterHandler("sensor/dht22_3/data", s.handleDHT221)
	s.MQTT.RegisterHandler("sensor/dht22_4/data", s.handleDHT221)
	s.MQTT.RegisterHandler("sensor/door/data", s.handleDoorSensor)
	s.MQTT.RegisterHandler("sensor/ds18b20/data", s.handleDs18b20)
	s.MQTT.RegisterHandler("sensor/pir/data", s.handlePIR)
	s.MQTT.RegisterHandler("sensor/leak_1/data", s.handleLeak1)
	s.MQTT.RegisterHandler("sensor/leak_2/data", s.handleLeak2)
	return s
}

func (s *Service) Start() error {
	return s.MQTT.Connect()
}
