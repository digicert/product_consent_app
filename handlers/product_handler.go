// product_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/digicert/product-consent-app/models"
	"github.com/digicert/product-consent-app/repository"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductRepo repository.ProductRepository
}

func NewProductHandler(productRepo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepo: productRepo,
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := ph.ProductRepo.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id, "message": "Product created successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse product id from url parameter

	vars := mux.Vars(r)
	productID := vars["id"]
	name := vars["name"]

	product := models.Product{
		ID:   productID,
		Name: name,
	}

	// Update product in database
	id, err := ph.ProductRepo.UpdateProduct(product.ID, product.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update product: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]string{"id": id, "message": "Product updated successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Parse product id from url parameter
	vars := mux.Vars(r)
	productId := vars["id"]

	product := models.Product{
		ID: productId,
	}

	// Delete product from database
	id, err := ph.ProductRepo.DeleteProduct(product.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete product: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]string{"id": id, "message": "Product deleted successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id"]

	product := models.Product{
		ID: productId,
	}

	// Get product from database
	products, err := ph.ProductRepo.GetProductById(product.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get product: %v", err), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse product json : %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// GetAllProducts returns all products
func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	products, err := ph.ProductRepo.GetAllProducts(offset, pageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get products: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse products json : %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
