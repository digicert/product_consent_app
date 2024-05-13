package controllers

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

// ProductController represents the REST controller for products
type ProductController struct {
	ProductHandler *handlers.ProductHandler
}

// New ProductController intializes a new instance of ProductController.
func NewproductController(productHandler *handlers.ProductHandler) *ProductController {
	return &ProductController{
		ProductHandler: productHandler,
	}
}

// Register routes for the product controller
func (pc *ProductController) RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", pc.ProductHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}/name/{name}", pc.ProductHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", pc.ProductHandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", pc.ProductHandler.GetProduct).Methods("GET")
	router.HandleFunc("/products", pc.ProductHandler.GetAllProducts).Methods("GET")
}
