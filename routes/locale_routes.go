package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

func RegisterLocaleRoutes(router *mux.Router, localeHandler *handlers.LocaleHandler) {
	router.HandleFunc("/locales", localeHandler.CreateLocale).Methods("POST")
	router.HandleFunc("/locales/{id}", localeHandler.UpdateLocale).Methods("PUT")
}
