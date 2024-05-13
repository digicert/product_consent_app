package pdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

type PDFGenerator struct{}

type Config struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (pg *PDFGenerator) GeneratePDF(configFile string) ([]byte, error) {
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %v", err)
	}

	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file %v", err)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, config.Title)
	pdf.Ln(10)
	pdf.MultiCell(0, 10, config.Content, "", "", false)

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to output pdf %v", err)
	}
	return buf.Bytes(), nil

}
