package handler

import (
	"SmartHouseAPI/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateUser(ctxTimeout, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}
