package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
	"github.com/gorilla/mux"
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

	id, err := lh.LocaleRepository.CreateLocale(&locale)
	if err != nil {
		http.Error(w, "failed to create locale", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Locale created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// UpdateLocale updates an existing locale
func (lh *LocaleHandler) UpdateLocale(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to update an existing locale
	mux := mux.Vars(r)
	localeID := mux["id"]
	name := mux["name"]

	locale := models.Locale{
		ID:     localeID,
		Locale: name,
	}

	id, err := lh.LocaleRepository.UpdateLocale(locale.ID, locale.Locale)
	if err != nil {
		http.Error(w, "failed to update locale", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Locale updated successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (lh *LocaleHandler) DeleteLocale(w http.ResponseWriter, r *http.Request) {
	// logic to delete an existing locale
	mux := mux.Vars(r)
	localeID := mux["id"]

	locale := models.Locale{
		ID: localeID,
	}

	id, err := lh.LocaleRepository.DeleteLocale(locale.ID)
	if err != nil {
		http.Error(w, "failed to delete locale", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"id": id, "message": "Locale deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (lh *LocaleHandler) GetLocale(w http.ResponseWriter, r *http.Request) {
	mux := mux.Vars(r)
	localeId := mux["id"]

	locale := models.Locale{
		ID: localeId,
	}

	locales, err := lh.LocaleRepository.GetLocaleById(locale.ID)
	if err != nil {
		http.Error(w, "failed to get locale"+err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(locales)
	w.Header().Set("Content-Tpe", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (lh *LocaleHandler) GetAllLocales(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	products, err := lh.LocaleRepository.GetAllLocales(offset, pageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get locales: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse locales json : %v", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
