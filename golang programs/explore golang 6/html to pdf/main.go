package main

import (
	"fmt"
	"os"
	"os/exec"
)

// ConvertHTMLToPDF converts an HTML file to a PDF file using wkhtmltopdf
func ConvertHTMLToPDF(inputHTML, outputPDF string) error {
	// Path to the wkhtmltopdf executable
	cmdPath := "wkhtmltopdf"

	// Prepare the command arguments
	args := []string{inputHTML, outputPDF}

	// Execute the command
	cmd := exec.Command(cmdPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and capture any error
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to convert HTML to PDF: %w", err)
	}

	return nil
}

func main() {
	inputHTML := "index.html"
	outputPDF := "output.pdf"

	if err := ConvertHTMLToPDF(inputHTML, outputPDF); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully")
}
