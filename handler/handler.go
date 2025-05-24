package handler

import (
	"SmartHouseAPI/models"
	"context"
)

type Service interface {
	AddClimateControlRecord(ctx context.Context, record models.ClimateControl) error
	AddLightRecord(ctx context.Context, record models.Light) error
	GenerateJWTToken() (string, error)
	GetClimateControlRecord(ctx context.Context, id int) (*models.ClimateControl, error)
	ValidateJWTToken(tokenString string) (bool, error)
	UpdateUser(ctx context.Context, user models.User) error
	SetLight(ctx context.Context, lightID int, on bool, brightness int) error
	SetHeater(ctx context.Context, on bool, temperature float64) error
	SetCooler(ctx context.Context, on bool, temperature float64) error
	GetWaterLeakSensorRecord(ctx context.Context, id int) (*models.WaterLeakSensor, error)
	GetUser(ctx context.Context, username string) (*models.User, error)
	GetMotionSensorRecord(ctx context.Context, id int) (*models.MotionSensor, error)
	GetLightSensorRecord(ctx context.Context, id int) (*models.LightSensor, error)
	GetLightRecord(ctx context.Context, id int) (*models.Light, error)
	GetDoorSensorRecord(ctx context.Context, id int) (*models.DoorSensor, error)
	GetDHT22SensorRecord(ctx context.Context, id int) (*models.DHT22Sensor, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}
