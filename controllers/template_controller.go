package controllers

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

type TemplateController struct {
	ConsentTemplateHandler *handlers.ConsentTemplateHandler
}

func NewTemplateController(consentTemplateHandler *handlers.ConsentTemplateHandler) *TemplateController {
	return &TemplateController{
		ConsentTemplateHandler: consentTemplateHandler,
	}
}

func (tc *TemplateController) RegisterTemplateRoutes(router *mux.Router) {
	router.HandleFunc("/templates", tc.ConsentTemplateHandler.CreateConsentTemplate).Methods("POST")
	router.HandleFunc("/templates/{id}", tc.ConsentTemplateHandler.GetConsentTemplateByID).Methods("GET")
}
