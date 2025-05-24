package handler

import (
	"SmartHouseAPI/models"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) SetCooler(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var req models.SetClimateControl
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.service.SetCooler(timeoutCtx, req.IsOn, req.TargetTemperature)
	if err != nil {
		log.Printf("SetCooler: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "SetCooler"})
		return
	}
	err = h.service.AddClimateControlRecord(timeoutCtx, models.ClimateControl{
		Sensor: models.Sensor{
			Name: "cooler",
		},
		IsOn:       req.IsOn,
		TargetTemp: req.TargetTemperature,
	})
	if err != nil {
		log.Printf("AddClimateControl: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "AddClimateControl"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}
