package controllers

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

type ProductTemplateController struct {
	ProductTemplateHandler *handlers.ProductTemplateHandler
}

func NewProductTemplateController(productTemplateHandler *handlers.ProductTemplateHandler) *ProductTemplateController {
	return &ProductTemplateController{
		ProductTemplateHandler: productTemplateHandler,
	}
}

func (ptc *ProductTemplateController) RegisterProductTemplateRoutes(router *mux.Router) {
	// Register routes for the product template controller
	router.HandleFunc("/product/templates", ptc.ProductTemplateHandler.CreateProductTemplate).Methods("POST")
	router.HandleFunc("/product/templates/{id}", ptc.ProductTemplateHandler.GetProductTemplateByID).Methods("GET")
	router.HandleFunc("/product/{id}/templates/active", ptc.ProductTemplateHandler.GetActiveProductTemplatesByProductID).Methods("GET")
}
