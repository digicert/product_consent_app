package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
	"github.com/gorilla/mux"
)

type ClientConsentHandler struct {
	ClientConsentRepo repository.ClientConsentRepository
}

func NewClientConsentHandler(clientConsentRepo repository.ClientConsentRepository) *ClientConsentHandler {
	return &ClientConsentHandler{
		ClientConsentRepo: clientConsentRepo,
	}
}

// CreateClientConsent creates a new client consent
func (cch *ClientConsentHandler) CreateClientConsent(w http.ResponseWriter, r *http.Request) {
	var clientConsent models.ClientConsent
	err := json.NewDecoder(r.Body).Decode(&clientConsent)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	id, err := cch.ClientConsentRepo.CreateClientConsent(&clientConsent)
	if err != nil {
		http.Error(w, "failed to create client consent", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Client consent created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// Get ClientConsent by ID
func (cch *ClientConsentHandler) GetClientConsentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	clientConsent, err := cch.ClientConsentRepo.GetClientConsentByID(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get client consent by ID: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(clientConsent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
