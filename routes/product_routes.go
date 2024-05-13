package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers routes for product-related endpoints
func RegisterProductRoutes(router *mux.Router, productHandler *handlers.ProductHandler) {
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}/name/{name}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
}
