package template

import (
	"fmt"
	"strconv"
	"time"

	"go-service/internal/types"

	"github.com/jung-kurt/gofpdf"
)

type Pdf = gofpdf.Fpdf

func isNonEmpty(fields ...string) bool {
	for _, f := range fields {
		if f != "" {
			return true
		}
	}
	return false
}

func isNonZeroInt(fields ...int) bool {
	for _, f := range fields {
		if f != 0 {
			return true
		}
	}
	return false
}

func GenerateStudentReportPDF(student types.Student) (*Pdf, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetMargins(20, 20, 20)

	// Title
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "STUDENT INFORMATION REPORT")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, "Generated on: " + time.Now().UTC().Format("2006-01-02")) // UTC datetime, to avoid timezone issues
	pdf.Ln(12)

	// Personal Information
	if isNonEmpty(student.Name, student.Email, student.Phone, student.Gender, student.DOB) || isNonZeroInt(student.ID) {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 8, "Personal Information")
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 12)
		if student.Name != "" {
			pdf.CellFormat(50, 8, "Full Name:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.Name, "", 1, "L", false, 0, "")
		}
		if student.ID != 0 {
			pdf.CellFormat(50, 8, "Student ID:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, strconv.Itoa(student.ID), "", 1, "L", false, 0, "")
		}
		if student.Email != "" {
			pdf.CellFormat(50, 8, "Email Address:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.Email, "", 1, "L", false, 0, "")
		}
		if student.Phone != "" {
			pdf.CellFormat(50, 8, "Phone Number:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.Phone, "", 1, "L", false, 0, "")
		}
		if student.Gender != "" {
			pdf.CellFormat(50, 8, "Gender:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.Gender, "", 1, "L", false, 0, "")
		}
		if student.DOB != "" {
			pdf.CellFormat(50, 8, "Date of Birth:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.DOB, "", 1, "L", false, 0, "")
		}
		pdf.Ln(6)
	}

	// Academic Information
	if isNonEmpty(student.Class, student.Section, student.AdmissionDate) || isNonZeroInt(student.Roll) {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 8, "Academic Information")
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 12)
		if student.Class != "" {
			pdf.CellFormat(50, 8, "Class:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.Class, "", 1, "L", false, 0, "")
		}
		if student.Section != "" {
			pdf.CellFormat(50, 8, "Section:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.Section, "", 1, "L", false, 0, "")
		}
		if student.Roll != 0 {
			pdf.CellFormat(50, 8, "Roll Number:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, strconv.Itoa(student.Roll), "", 1, "L", false, 0, "")
		}
		if student.AdmissionDate != "" {
			pdf.CellFormat(50, 8, "Admission Date:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.AdmissionDate, "", 1, "L", false, 0, "")
		}
		pdf.Ln(6)
	}

	// Parent Information
	if isNonEmpty(student.FatherName, student.FatherPhone, student.MotherName, student.MotherPhone) {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 8, "Parent Information")
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 12)
		if student.FatherName != "" {
			pdf.CellFormat(50, 8, "Father's Name:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.FatherName, "", 1, "L", false, 0, "")
		}
		if student.FatherPhone != "" {
			pdf.CellFormat(50, 8, "Father's Phone:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.FatherPhone, "", 1, "L", false, 0, "")
		}
		if student.MotherName != "" {
			pdf.CellFormat(50, 8, "Mother's Name:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.MotherName, "", 1, "L", false, 0, "")
		}
		if student.MotherPhone != "" {
			pdf.CellFormat(50, 8, "Mother's Phone:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.MotherPhone, "", 1, "L", false, 0, "")
		}
		pdf.Ln(6)
	}

	// Guardian Information
	if isNonEmpty(student.GuardianName, student.GuardianPhone, student.RelationOfGuardian) {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 8, "Guardian Information")
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 12)
		if student.GuardianName != "" {
			pdf.CellFormat(50, 8, "Guardian's Name:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.GuardianName, "", 1, "L", false, 0, "")
		}
		if student.GuardianPhone != "" {
			pdf.CellFormat(50, 8, "Guardian's Phone:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.GuardianPhone, "", 1, "L", false, 0, "")
		}
		if student.RelationOfGuardian != "" {
			pdf.CellFormat(50, 8, "Relation:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.RelationOfGuardian, "", 1, "L", false, 0, "")
		}
		pdf.Ln(6)
	}

	// Address Information
	if isNonEmpty(student.CurrentAddress, student.PermanentAddress) {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 8, "Address Information")
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 12)
		if student.CurrentAddress != "" {
			pdf.CellFormat(50, 8, "Current Address:", "", 0, "L", false, 0, "")
			pdf.MultiCell(0, 8, student.CurrentAddress, "", "L", false)
		}
		if student.PermanentAddress != "" {
			pdf.CellFormat(50, 8, "Permanent Address:", "", 0, "L", false, 0, "")
			pdf.MultiCell(0, 8, student.PermanentAddress, "", "L", false)
		}
		pdf.Ln(6)
	}

	// Additional Information
	if isNonEmpty(student.ReporterName) || student.SystemAccess {
		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 8, "Additional Information")
		pdf.Ln(10)
		pdf.SetFont("Arial", "", 12)
		if student.ReporterName != "" {
			pdf.CellFormat(50, 8, "Reporter:", "", 0, "L", false, 0, "")
			pdf.CellFormat(0, 8, student.ReporterName, "", 1, "L", false, 0, "")
		}
		pdf.CellFormat(50, 8, "System Access:", "", 0, "L", false, 0, "")
		pdf.CellFormat(0, 8, fmt.Sprintf("%t", student.SystemAccess), "", 1, "L", false, 0, "")
		pdf.Ln(6)
	}

	pdf.Ln(10)
	pdf.SetFont("Arial", "I", 10)
	pdf.Cell(0, 8, "This is an official document generated by the School Management System.")
	pdf.Ln(6)
	pdf.Cell(0, 8, "For any queries, please contact the school administration.")

	return pdf, nil
} 