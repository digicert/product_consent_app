package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
	"github.com/gorilla/mux"
)

type ProductTemplateHandler struct {
	ProductTemplateRepo repository.ProductTemplateRepository
}

func NewProductTemplateHandler(productTemplateRepo repository.ProductTemplateRepository) *ProductTemplateHandler {
	return &ProductTemplateHandler{
		ProductTemplateRepo: productTemplateRepo,
	}
}

func (pth *ProductTemplateHandler) CreateProductTemplate(w http.ResponseWriter, r *http.Request) {
	// CreateProductTemplate creates a new product template

	var productTemplate models.ProductTemplate
	err := json.NewDecoder(r.Body).Decode(&productTemplate)

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	id, err := pth.ProductTemplateRepo.CreateProductTemplate(&productTemplate)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create product template: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Product template created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (pth *ProductTemplateHandler) GetProductTemplateByID(w http.ResponseWriter, r *http.Request) {
	// GetProductTemplateByID retrieves a product template by its ID

	vars := mux.Vars(r)
	ID := vars["id"]

	productTemplate, err := pth.ProductTemplateRepo.GetProductTemplateByID(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get product template by ID: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(productTemplate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (pth *ProductTemplateHandler) GetActiveProductTemplatesByProductID(w http.ResponseWriter, r *http.Request) {
	// GetActiveProductTemplatesByProductID retrieves all active product templates by product ID

	vars := mux.Vars(r)
	productID := vars["id"]

	productTemplates, err := pth.ProductTemplateRepo.GetActiveProductTemplatesByProductID(productID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get active product templates by product ID: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, _ := json.Marshal(productTemplates)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
