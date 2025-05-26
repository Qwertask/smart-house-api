package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	key    = "Authorization"
	prefix = "Bearer "
)

func (h *Handler) JWTCheckMiddleware(c *gin.Context) {
	ctx := c.Request.Context()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	authHeader := c.GetHeader(key)
	if authHeader == "" || !strings.HasPrefix(authHeader, prefix) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found in context"})
		c.Abort()
		return
	}
	token := strings.TrimPrefix(authHeader, prefix)
	tokenInvalid, err := h.cache.Get(timeoutCtx, token)
	if err != nil {
		log.Println(err)
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			c.Abort()
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on server side"})
		c.Abort()
		return
	}
	if tokenInvalid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	tokenValid, err := h.service.ValidateJWTToken(token)
	if err != nil {
		log.Println(err)
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	if !tokenValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
