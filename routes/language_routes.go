package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

func RegisterLanguageRoutes(router *mux.Router, languageHandler *handlers.LanguageHandler) {
	router.HandleFunc("/languages", languageHandler.CreateLanguage).Methods("POST")
	router.HandleFunc("/languages/{id}", languageHandler.UpdateLanguage).Methods("PUT")
	router.HandleFunc("/languages/{language_id}/locales/{locale_id}", languageHandler.LinkLanguageWithLocale).Methods("POST")
}
