package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go-service/internal/config"
	"go-service/internal/template"
	"go-service/internal/types"

	"github.com/gin-gonic/gin"
)

type StudentsService struct{}

func NewStudentsService() *StudentsService {
	return &StudentsService{}
}

func (s *StudentsService) GetStudentReportPDF(c *gin.Context, studentID string) (*template.Pdf, error) {
	// Fetch student data from backend
	cfg := config.Load()
	backendURL := fmt.Sprintf("%s/api/v1/students/%s", cfg.ApiURL, studentID)

	req, err := http.NewRequest("GET", backendURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create backend request: %w", err)
	}

	// Forward Cookie header
	if cookie := c.Request.Header.Get("Cookie"); cookie != "" {
		req.Header.Set("Cookie", cookie)
	}

	// Forward x-csrf-token header
	if csrf := c.Request.Header.Get("x-csrf-token"); csrf != "" {
		req.Header.Set("x-csrf-token", csrf)
	}

	client := &http.Client{}
	backendResponse, err := client.Do(req)

	// Handle error
	if err != nil {
		return nil, fmt.Errorf("failed to get student report from backend: %w", err)
	}
	defer backendResponse.Body.Close()

	if backendResponse.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(backendResponse.Body)
		msg := strings.TrimSpace(string(body))
		if msg == "" {
			msg = fmt.Sprintf("backend returned status %d", backendResponse.StatusCode)
		}
		return nil, errors.New(msg)
	}

	// Read and parse student data
	body, err := io.ReadAll(backendResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read backend response: %w", err)
	}

	var student types.Student
	if err := json.Unmarshal(body, &student); err != nil {
		return nil, fmt.Errorf("failed to parse student data: %w", err)
	}

	// Generate PDF report
	return template.GenerateStudentReportPDF(student)
}
