package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

func RegisterLocaleRoutes(router *mux.Router, localeHandler *handlers.LocaleHandler) {
	router.HandleFunc("/locales", localeHandler.CreateLocale).Methods("POST")
	router.HandleFunc("/locales/{id}/name/{name}", localeHandler.UpdateLocale).Methods("PUT")
	router.HandleFunc("/locales/{id}", localeHandler.DeleteLocale).Methods("DELETE")
	router.HandleFunc("/locales/{id}", localeHandler.GetLocale).Methods("GET")
	router.HandleFunc("/locales", localeHandler.GetAllLocales).Methods("GET")
}
