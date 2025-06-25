package routes

import (
	"go-service/internal/handler"
	"go-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	healthHandler *handler.HealthHandler,
	reportHandler *handler.ReportHandler,
) *gin.Engine {
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORS())
	r.Use(middleware.Logging())

	// Health check
	r.GET("/health", healthHandler.HealthCheck)

	// Root endpoint
	r.GET("/", healthHandler.ServiceStatus)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("/:id/report", reportHandler.GetStudentReport)
		}
	}

	return r
}
