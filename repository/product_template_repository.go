package repository

import (
	"database/sql"
	"fmt"

	"github.com/digicert/product-consent-app/models"
	"github.com/google/uuid"
)

type ProductTemplateRepository struct {
	DB *sql.DB
}

func NewProductTemplateRepository(db *sql.DB) *ProductTemplateRepository {
	return &ProductTemplateRepository{
		DB: db,
	}
}

func (ptr *ProductTemplateRepository) CreateProductTemplate(productTemplate *models.ProductTemplate) (string, error) {
	ID := uuid.New().String()
	_, err := ptr.DB.Exec("INSERT INTO product_template (id, product_id, consent_template_id, active) VALUES (?, ?, ?, ?)", ID, productTemplate.ProductID, productTemplate.ConsentTemplateID, productTemplate.Active)
	if err != nil {
		return "", fmt.Errorf("failed to create product template: %v", err)
	}
	return ID, nil
}

func (ptr *ProductTemplateRepository) GetProductTemplateByID(ID string) (*models.ProductTemplate, error) {
	var productTemplate models.ProductTemplate
	err := ptr.DB.QueryRow("SELECT id, product_id, consent_template_id, active FROM product_template WHERE id = ?", ID).Scan(&productTemplate.ID, &productTemplate.ProductID, &productTemplate.ConsentTemplateID, &productTemplate.Active)
	if err != nil {
		return nil, fmt.Errorf("failed to get product template by ID: %v", err)
	}
	return &productTemplate, nil
}
