// handlers/consent_template_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/pdf"
	"github.com/digicert/product-consent-app/repository"
	"github.com/gorilla/mux"
)

// ConsentTemplateHandler represents the REST handler for consent templates.
type ConsentTemplateHandler struct {
	ConsentTemplateRepo repository.ConsentTemplateRepository
}

// NewConsentTemplateHandler initializes a new instance of ConsentTemplateHandler.
func NewConsentTemplateHandler(consentTemplateRepo repository.ConsentTemplateRepository) *ConsentTemplateHandler {
	return &ConsentTemplateHandler{
		ConsentTemplateRepo: consentTemplateRepo,
	}
}

// CreateConsentTemplate handles the creation of a new consent template.
func (cth *ConsentTemplateHandler) CreateConsentTemplate(w http.ResponseWriter, r *http.Request) {
	var consentTemplate models.ConsentTemplate

	err := json.NewDecoder(r.Body).Decode(&consentTemplate)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	pdfContent, err := pdf.LoadPDFContent("config/template_config.json", consentTemplate.TemplateName)

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to load PDF content: %v", err), http.StatusInternalServerError)
		return
	}
	consentTemplate.TemplatePDF = pdfContent

	// Call repository method to create consent template
	id, err := cth.ConsentTemplateRepo.CreateConsentTemplate(&consentTemplate)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create consent template: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	response := map[string]string{"id": id, "message": "Consent template created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// GetConsentTemplateByID retrieves a consent template by its ID.
func (cth *ConsentTemplateHandler) GetConsentTemplateByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from request path parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Call repository method to get consent template by ID
	consentTemplate, err := cth.ConsentTemplateRepo.GetConsentTemplateByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get consent template: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with consent template
	jsonResponse, err := json.Marshal(consentTemplate)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal JSON: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Implement other handler methods for updating, deleting, and retrieving all consent templates...
