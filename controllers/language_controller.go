package controllers

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

// LanguageController is a controller for the Language model
type LanguageController struct {
	LanguageHandler *handlers.LanguageHandler
}

// NewLanguageController initializes a new instance of LanguageController
func NewLanguageController(languageHandler *handlers.LanguageHandler) *LanguageController {
	return &LanguageController{
		LanguageHandler: languageHandler,
	}
}

// RegisterLanguageRoutes registers routes for the language controller
func (lc *LanguageController) RegisterLanguageRoutes(router *mux.Router) {
	router.HandleFunc("/languages", lc.LanguageHandler.CreateLanguage).Methods("POST")
	router.HandleFunc("/languages/{id}/name/{name}", lc.LanguageHandler.UpdateLanguage).Methods("PUT")
	router.HandleFunc("/languages/{id}", lc.LanguageHandler.DeleteLanguage).Methods("DELETE")
	router.HandleFunc("/languages/{id}", lc.LanguageHandler.GetLanguageByID).Methods("GET")
	router.HandleFunc("/languages", lc.LanguageHandler.GetAllLanguages).Methods("GET")
	// link language to locale
	router.HandleFunc("/languages/{language_id}/locales/{locale_id}", lc.LanguageHandler.LinkLanguageWithLocale).Methods("POST")
	// unlink language from locale
	router.HandleFunc("/languages/{language_id}/locales/{locale_id}", lc.LanguageHandler.UnlinkLanguageWithLocale).Methods("DELETE")
	// get all linked languages for a locale
	router.HandleFunc("/locales/{locale_id}/languages", lc.LanguageHandler.GetLinkedLanguagesByLocaleID).Methods("GET")

}
