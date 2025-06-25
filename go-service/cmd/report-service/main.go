package main

import (
	"log"

	"go-service/internal/config"
	"go-service/internal/handler"
	"go-service/internal/routes"
	"go-service/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services
	reportService := service.NewReportService()

	// Initialize handlers
	healthHandler := handler.NewHealthHandler()
	reportHandler := handler.NewReportHandler(reportService)

	// Setup routes
	r := routes.SetupRoutes(healthHandler, reportHandler)

	// Log startup message
	log.Printf("🚀 Go microservice starting on port %s", cfg.Port)
	log.Printf("📡 Available endpoints:")
	log.Printf("   GET / - Service status")
	log.Printf("   GET /health - Health check")
	log.Printf("   GET /api/v1/students/:id/report - Student report (PDF)")

	// Start the server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
