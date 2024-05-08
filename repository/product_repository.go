package repository

import (
	"database/sql"
	"fmt"

	"github.com/digicert/product-consent-app/models"
	"github.com/google/uuid"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

// CreateProduct creates a new product in the database
func (pr *ProductRepository) CreateProduct(product *models.Product) (string, error) {
	product.ID = uuid.New().String()
	_, err := pr.DB.Exec("INSERT INTO product (id, name) VALUES (?, ?)", product.ID, product.Name)
	if err != nil {
		return "", fmt.Errorf("failed to create product: %v", err)
	}
	return product.ID, nil
}

// UpdateProduct updates an existing product in the database
func (pr *ProductRepository) UpdateProduct(ID string, name string) (string, error) {
	result, err := pr.DB.Exec("UPDATE product SET name = ? WHERE id = ?", name, ID)
	if err != nil {
		return "", fmt.Errorf("failed to update product: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, product not updated")
	}
	return ID, err
}

func (pr *ProductRepository) DeleteProduct(ID string) (string, error) {
	result, err := pr.DB.Exec("DELETE FROM product WHERE id = ?", ID)
	if err != nil {
		return "", fmt.Errorf("failed to delete product: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, product not deleted")
	}
	return ID, err
}
