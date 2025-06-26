package handler

import (
	"log"
	"net/http"

	"go-service/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentsHandler struct {
	studentsService *service.StudentsService
}

func NewStudentsHandler(studentsService *service.StudentsService) *StudentsHandler {
	return &StudentsHandler{
		studentsService: studentsService,
	}
}

func (h *StudentsHandler) GetStudentReport(c *gin.Context) {
	studentID := c.Param("id")

	// Generate student report PDF
	pdf, err := h.studentsService.GetStudentReportPDF(c, studentID)
	if err != nil {
		log.Printf("Failed to get/generate student report: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=student_report_"+studentID+".pdf")
	c.Status(http.StatusOK)

	if err := pdf.Output(c.Writer); err != nil {
		log.Printf("Failed to output PDF: %v", err)
		c.String(http.StatusInternalServerError, "Failed to output PDF")
	}
}
