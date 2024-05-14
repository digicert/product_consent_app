package repository

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/digicert/product-consent-app/models"
	"github.com/google/uuid"
)

type ClientConsentRepository struct {
	DB *sql.DB
}

func NewClientConsentRepository(db *sql.DB) *ClientConsentRepository {
	return &ClientConsentRepository{
		DB: db,
	}
}

func (ccr *ClientConsentRepository) CreateClientConsent(clientConsent *models.ClientConsent) (string, error) {
	ID := uuid.New().String()
	// date and time
	Datetime := time.Now()

	_, err := ccr.DB.Exec("INSERT INTO client_consent (id, product_template_id, individual_id, date, optout_reason) VALUES (?, ?, ?, ?, ?)", ID, clientConsent.ProductTemplateID, clientConsent.IndividualID, Datetime, clientConsent.OptoutReason)
	if err != nil {
		return "", fmt.Errorf("failed to create client consent: %v", err)
	}
	return ID, nil
}

func (ccr *ClientConsentRepository) GetClientConsentByID(ID string) (*models.ClientConsent, error) {
	var clientConsent models.ClientConsent
	err := ccr.DB.QueryRow("SELECT id, product_template_id, individual_id, date, optout_reason FROM client_consent WHERE id = ?", ID).Scan(&clientConsent.ID, &clientConsent.ProductTemplateID, &clientConsent.IndividualID, &clientConsent.Date, &clientConsent.OptoutReason)
	if err != nil {
		return nil, fmt.Errorf("failed to get client consent by ID: %v", err)
	}
	return &clientConsent, nil
}
