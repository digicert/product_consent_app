package main

import (
	"log"
	"net/http"

	"github.com/digicert/product-consent-app/db"
	"github.com/digicert/product-consent-app/handlers"
	"github.com/digicert/product-consent-app/repository"
	"github.com/digicert/product-consent-app/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	db.InitDB()
	defer db.CloseDB()

	// Initialize router
	router := mux.NewRouter()

	// Initialize product repository
	productRepo := repository.NewProductRepository(db.DB)
	languageRepo := repository.NewLanguageRepository(db.DB)
	localeRepo := repository.NewLocaleRepository(db.DB)
	templateRepo := repository.NewConsentTemplateRepository(db.DB)
	productTemplateRepo := repository.NewProductTemplateRepository(db.DB)
	clientConsentRepo := repository.NewClientConsentRepository(db.DB)

	// Initialize product handler with the product repository
	productHandler := handlers.NewProductHandler(*productRepo)
	languageHandler := handlers.NewLanguageHandler(*languageRepo)
	localeHandler := handlers.NewLocaleHandler(*localeRepo)
	templateHandler := handlers.NewConsentTemplateHandler(*templateRepo)
	productTemplateHandler := handlers.NewProductTemplateHandler(*productTemplateRepo)
	clientConsentHandler := handlers.NewClientConsentHandler(*clientConsentRepo)

	// Register product routes
	routes.RegisterProductRoutes(router, productHandler)
	routes.RegisterLanguageRoutes(router, languageHandler)
	routes.RegisterLocaleRoutes(router, localeHandler)
	routes.RegisterTemplateRoutes(router, templateHandler)
	routes.RegisterProductTemplateRoutes(router, productTemplateHandler)
	routes.RegisterClientConsentRoutes(router, clientConsentHandler)

	// Start server
	log.Println("Server started on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))

}
