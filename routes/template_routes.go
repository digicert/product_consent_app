package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

// RegisterTemplateRoutes registers routes for the template controller
func RegisterTemplateRoutes(router *mux.Router, templateHandler *handlers.ConsentTemplateHandler) {
	router.HandleFunc("/templates", templateHandler.CreateConsentTemplate).Methods("POST")
	router.HandleFunc("/templates/{id}", templateHandler.GetConsentTemplateByID).Methods("GET")
}
