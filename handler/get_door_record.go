package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) GetDoorRecord(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	record, err := h.service.GetDoorSensorRecord(timeoutCtx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": record.Name, "is_open": record.IsDoorOpen, "timestamp": record.Timestamp.Format(time.RFC3339)})
	return
}
