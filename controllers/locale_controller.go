package controllers

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

// LocaleController represents the REST controller for locales
type LocaleController struct {
	LocaleHandler *handlers.LocaleHandler
}

// NewLocaleController initializes a new instance of LocaleController
func NewLocaleController(localeHandler *handlers.LocaleHandler) *LocaleController {
	return &LocaleController{
		LocaleHandler: localeHandler,
	}
}

// RegisterLocaleRoutes registers routes for the locale controller
func (lc *LocaleController) RegisterLocaleRoutes(router *mux.Router) {
	router.HandleFunc("/locales", lc.LocaleHandler.CreateLocale).Methods("POST")
	router.HandleFunc("/locales/{id}/name/{name}", lc.LocaleHandler.UpdateLocale).Methods("PUT")
	router.HandleFunc("/locales/{id}", lc.LocaleHandler.DeleteLocale).Methods("DELETE")
}
