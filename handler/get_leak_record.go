package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetLeakRecord(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	sensorID, err := strconv.Atoi(c.Param("leak_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record, err := h.service.GetWaterLeakSensorRecord(timeoutCtx, sensorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": record.Name, "leak": record.IsLeaking, "timestamp": record.Timestamp.Format(time.RFC3339)})
	return
}
