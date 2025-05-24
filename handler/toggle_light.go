package handler

import (
	"SmartHouseAPI/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) ToggleLight(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var req models.SetLight
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.service.SetLight(timeoutCtx, req.ID, req.IsOn, req.Brightness)
	if err != nil {
		log.Printf("SetLight: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "SetLight failed"})
		return
	}
	err = h.service.AddLightRecord(timeoutCtx, models.Light{
		Sensor: models.Sensor{
			Name: fmt.Sprintf("light_%d", req.ID),
		},
		IsOn:       req.IsOn,
		Brightness: req.Brightness,
	})
	if err != nil {
		log.Printf("Toggle light: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Toggle light failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}
