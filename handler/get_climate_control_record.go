package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) GetClimateControlRecord(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	device := c.Query("device")
	record, err := h.service.GetClimateControlRecord(timeoutCtx, device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": record.Name, "is_on": record.IsOn, "target_temp": record.TargetTemp})
	return
}
