package handler

import (
	"log"
	"net/http"

	"go-service/internal/service"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService *service.ReportService
}

func NewReportHandler(reportService *service.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

func (h *ReportHandler) GetStudentReport(c *gin.Context) {
	studentID := c.Param("id")

	pdf, err := h.reportService.GenerateStudentReport(studentID)
	if err != nil {
		log.Printf("Failed to generate PDF: %v", err)
		c.String(http.StatusInternalServerError, "Failed to generate PDF")
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=student_report_"+studentID+".pdf")
	c.Status(http.StatusOK)

	if err := pdf.Output(c.Writer); err != nil {
		log.Printf("Failed to output PDF: %v", err)
		c.String(http.StatusInternalServerError, "Failed to output PDF")
	}
}
