package main

import (
	"fmt"
	"strings"

	"github.com/go-pdf/fpdf"
)

func main() {

	var formatRect = func(pdf *fpdf.Fpdf, font string, style string, size float64, height float64, text string, align string) {
		pdf.SetAutoPageBreak(false, 0)
		pdf.SetXY(20, 20)
		pdf.SetFont(font, style, size)
		pdf.CellFormat(170, height, text, "", 0, align, false, 0, "")
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	//Image, X-Axis, Y-Axis, Width, height
	pdf.Image("./Images/GoogleLogo.png", 72, 12, 70, 20, false, "", 0, "")
	pdf.Image("./Images/ProfilePhoto.png", 150, 75, 30, 30, false, "", 0, "")
	pdf.Image("./Images/QRCode.png", 85, 200, 40, 40, false, "", 0, "")

	// Add Text in rectangle alignment
	formatRect(pdf, "Arial", "B", 18, 45, "New Certificate", "C")
	formatRect(pdf, "Arial", "B", 11, 128, "This is to certify that:", "L")
	formatRect(pdf, "Arial", "B", 11, 143, "Has been found buly qualified to be in charge of:", "L")
	formatRect(pdf, "Arial", "", 12, 143, GetStringWithLeftSpace("Flag Stat", 78), "L")

	// first
	formatRect(pdf, "Arial", "B", 11, 200, "Issued on: ", "L")
	formatRect(pdf, "Arial", "", 12, 200, GetStringWithLeftSpace("07-01-2022", 17), "L")
	formatRect(pdf, "Arial", "B", 11, 200, GetStringWithRightSpace("Valid untill:", 20), "R")
	formatRect(pdf, "Arial", "", 11, 200, GetStringWithLeftSpace("07-01-2023", 20), "R")

	// second
	formatRect(pdf, "Arial", "B", 11, 219, "Certificate no. ", "L")
	formatRect(pdf, "Arial", "", 12, 219, GetStringWithLeftSpace("4014c601bcaf-4290-a878", 23), "L")
	formatRect(pdf, "Arial", "B", 11, 219, GetStringWithRightSpace("Course ID:", 44), "R")
	formatRect(pdf, "Arial", "", 11, 219, GetStringWithLeftSpace("21b93553-a0db-4018-bf44", 44), "R")

	// third
	formatRect(pdf, "Arial", "B", 11, 238, "Passport number of the certificate holder ", "L")
	formatRect(pdf, "Arial", "", 12, 238, GetStringWithLeftSpace("*****38sd", 66), "L")
	formatRect(pdf, "Arial", "B", 11, 238, GetStringWithRightSpace("Date of birth:", 20), "R")
	formatRect(pdf, "Arial", "", 11, 238, GetStringWithLeftSpace("07-01-2022", 20), "R")

	formatRect(pdf, "Arial", "", 12, 335, "The issuers organisation is approved by XYZ", "C")
	formatRect(pdf, "Arial", "B", 16, 353, "Certificate QR:", "C")

	err := pdf.OutputFileAndClose("Certificate.pdf")
	if err != nil {
		fmt.Println("Error to create PDF")
	}
}

func GetStringWithLeftSpace(str string, space int) string {
	return (strings.Repeat(" ", space) + str)
}
func GetStringWithRightSpace(str string, space int) string {
	return (str + strings.Repeat(" ", space))
}
