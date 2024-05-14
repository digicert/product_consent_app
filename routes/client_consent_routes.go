package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

// RegisterClientConsentRoutes registers routes for the client consent controller
func RegisterClientConsentRoutes(router *mux.Router, clientConsentHandler *handlers.ClientConsentHandler) {
	router.HandleFunc("/client/consents", clientConsentHandler.CreateClientConsent).Methods("POST")
	router.HandleFunc("/client/consents/{id}", clientConsentHandler.GetClientConsentByID).Methods("GET")
}
