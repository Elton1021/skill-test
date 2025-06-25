package service

import (
	"github.com/jung-kurt/gofpdf"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) GenerateStudentReport(studentID string) (*gofpdf.Fpdf, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 24)
	pdf.Cell(40, 10, "Hi "+studentID+"!")

	return pdf, nil
}
