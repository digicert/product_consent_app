package controllers

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

type ClientConsentController struct {
	ClientConsentHandler *handlers.ClientConsentHandler
}

func NewClientConsentController(clientConsentHandler *handlers.ClientConsentHandler) *ClientConsentController {
	return &ClientConsentController{
		ClientConsentHandler: clientConsentHandler,
	}
}

func (ccc *ClientConsentController) RegisterClientConsentRoutes(router *mux.Router) {
	router.HandleFunc("/client/consents", ccc.ClientConsentHandler.CreateClientConsent).Methods("POST")
	router.HandleFunc("/client/consents/{id}", ccc.ClientConsentHandler.GetClientConsentByID).Methods("GET")
}
