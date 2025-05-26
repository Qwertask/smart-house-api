package handler

import (
	"SmartHouseAPI/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	var refreshToken models.LogoutRequest
	if err := c.BindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.cache.Set(ctxTimeout, accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.cache.Set(ctxTimeout, refreshToken.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
