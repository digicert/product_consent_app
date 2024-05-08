package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
	"github.com/gorilla/mux"
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
		http.Error(w, "failed to create language "+err.Error(), http.StatusInternalServerError)
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
	vars := mux.Vars(r)
	languageID := vars["id"]
	name := vars["name"]

	link := models.Language{
		ID:       languageID,
		Language: name,
	}

	id, err := lh.LanguageRepo.UpdateLanguage(link.ID, link.Language)
	if err != nil {
		http.Error(w, "failed to update language "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Language updated successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (lh *LanguageHandler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to delete an existing language
	vars := mux.Vars(r)
	languageID := vars["id"]

	link := models.Language{
		ID: languageID,
	}

	id, err := lh.LanguageRepo.DeleteLanguage(link.ID)
	if err != nil {
		http.Error(w, "failed to delete language "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Language deleted successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Link Language with Locale
func (lh *LanguageHandler) LinkLanguageWithLocale(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to link a language with a locale
	vars := mux.Vars(r)
	localeID := vars["locale_id"]
	languageID := vars["language_id"]
	fmt.Printf("Linking language with locale from params: LocaleID=%s, LanguageID=%s\n", localeID, languageID)

	link := models.LocaleLanguage{
		LocaleID:   localeID,
		LanguageID: languageID,
	}

	id, err := lh.LanguageRepo.LinkLanguageWithLocale(link.ID, link.LocaleID, link.LanguageID)
	if err != nil {
		http.Error(w, "failed to link language with locale "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Language linked with locale successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
