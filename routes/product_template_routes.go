package routes

import (
	"github.com/digicert/product-consent-app/handlers"
	"github.com/gorilla/mux"
)

func RegisterProductTemplateRoutes(router *mux.Router, productTemplateHandler *handlers.ProductTemplateHandler) {
	router.HandleFunc("/product/templates", productTemplateHandler.CreateProductTemplate).Methods("POST")
	router.HandleFunc("/product/templates/{id}", productTemplateHandler.GetProductTemplateByID).Methods("GET")
	router.HandleFunc("/product/{id}/templates/active", productTemplateHandler.GetActiveProductTemplatesByProductID).Methods("GET")
}
