package repository

import (
	"database/sql"
	"fmt"

	"github.com/digicert/product-consent-app/models"
	"github.com/google/uuid"
)

type ConsentTemplateRepository struct {
	DB *sql.DB
}

func NewConsentTemplateRepository(db *sql.DB) *ConsentTemplateRepository {
	return &ConsentTemplateRepository{
		DB: db,
	}
}

func (ctr *ConsentTemplateRepository) CreateConsentTemplate(consentTemplate *models.ConsentTemplate) (string, error) {
	ID := uuid.New().String()
	_, err := ctr.DB.Exec("INSERT INTO consent_template (id, locale_language_id, template_pdf) VALUES (?, ?, ?)", ID, consentTemplate.LocaleLanguageID, consentTemplate.TemplatePDF)
	if err != nil {
		return "", fmt.Errorf("failed to create consent template: %v", err)
	}
	return ID, nil
}

// GetConsentTemplateByID retrieves a consent template by its ID
func (ctr *ConsentTemplateRepository) GetConsentTemplateByID(ID string) (*models.ConsentTemplate, error) {
	var consentTemplate models.ConsentTemplate
	err := ctr.DB.QueryRow("SELECT id, locale_language_id, template_pdf FROM consent_template WHERE id = ?", ID).Scan(&consentTemplate.ID, &consentTemplate.LocaleLanguageID, &consentTemplate.TemplatePDF)
	if err != nil {
		return nil, fmt.Errorf("failed to get consent template by ID: %v", err)
	}
	return &consentTemplate, nil
}
