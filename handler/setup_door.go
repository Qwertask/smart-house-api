package handler

import (
	"SmartHouseAPI/models"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	openDoor  = "open"
	closeDoor = "close"
)

// post
func (h *Handler) SetupDoor(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var req models.SetupDoorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if req.Setup != openDoor && req.Setup != closeDoor {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	err := h.service.AddDoorSensorRecord(timeoutCtx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
