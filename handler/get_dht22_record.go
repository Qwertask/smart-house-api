package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetDHT22Record(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	sensorId, err := strconv.Atoi(c.Query("dht22id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record, err := h.service.GetDHT22SensorRecord(timeoutCtx, sensorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": record.Name, "temperature": record.Temperature, "humidity": record.Humidity, "timestamp": record.Timestamp.Format(time.RFC3339)})
	return
}
