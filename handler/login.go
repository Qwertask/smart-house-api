package handler

import (
	"SmartHouseAPI/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	ctxTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var req models.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, err := h.service.Login(ctxTimeout, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
