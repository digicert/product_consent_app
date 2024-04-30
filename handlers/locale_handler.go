package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
)

// LocaleHandler represents the handler for operating on locales
type LocaleHandler struct {
	LocaleRepository repository.LocaleRepository
}

// NewLocaleHandler initializes a new instance of LocaleHandler
func NewLocaleHandler(localeRepository repository.LocaleRepository) *LocaleHandler {
	return &LocaleHandler{
		LocaleRepository: localeRepository,
	}
}

// CreateLocale creates a new locale
func (lh *LocaleHandler) CreateLocale(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to create a new locale
	var locale models.Locale
	err := json.NewDecoder(r.Body).Decode(&locale)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	err = lh.LocaleRepository.CreateLocale(&locale)
	if err != nil {
		http.Error(w, "failed to create locale", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Locale created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// UpdateLocale updates an existing locale
func (lh *LocaleHandler) UpdateLocale(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to update an existing locale
	var locale models.Locale
	err := json.NewDecoder(r.Body).Decode(&locale)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	err = lh.LocaleRepository.UpdateLocale(&locale)
	if err != nil {
		http.Error(w, "failed to update locale", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Locale updated successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
