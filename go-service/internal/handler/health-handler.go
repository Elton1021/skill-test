package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "go-service",
	})
}

func (h *HealthHandler) ServiceStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Go Microservice is running!",
		"service": "go-service",
	})
}
