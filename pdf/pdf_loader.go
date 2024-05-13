package pdf

import (
	"encoding/json"
	"fmt"
	"os"
)

type PDFConfig struct {
	Templates map[string]string `json:"templates"`
}

func LoadPDFContent(configFilePath string, templateType string) ([]byte, error) {
	// Open the config file
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	// Decode the JSON config file
	var config PDFConfig
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}
	// Get the content for the specified template type
	pdfContent, ok := config.Templates[templateType]
	if !ok {
		return nil, fmt.Errorf("template type not found: %s", templateType)
	}
	return []byte(pdfContent), nil
}
