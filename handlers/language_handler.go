package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
)

// LanguageHandler represents the handler for managing languages.
type LanguageHandler struct {
	LanguageRepo repository.LanguageRepository
}

func NewLanguageHandler(languageRepo repository.LanguageRepository) *LanguageHandler {
	return &LanguageHandler{
		LanguageRepo: languageRepo,
	}
}

// CreateLanguage creates a new language.
func (lh *LanguageHandler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	var language models.Language

	err := json.NewDecoder(r.Body).Decode(&language)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	id, err := lh.LanguageRepo.CreateLanguage(&language)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Language created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// UpdateLanguage updates an existing language.
func (lh *LanguageHandler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to update an existing language
	var language models.Language
	err := json.NewDecoder(r.Body).Decode(&language)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	id, err := lh.LanguageRepo.UpdateLanguage(&language)
	if err != nil {
		http.Error(w, "failed to update language", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Language updated successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Link Language with Locale
func (lh *LanguageHandler) LinkLanguageWithLocale(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to link a language with a locale
	var link models.LocaleLanguage
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	id, err := lh.LanguageRepo.LinkLanguageWithLocale(link.LanguageID, link.LocaleID)
	if err != nil {
		http.Error(w, "failed to link language with locale", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Language linked with locale successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
