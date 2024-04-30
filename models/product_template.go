package models

type ProductTemplate struct {
	ID                string `json:"id"`
	ProductID         string `json:"product_id"`
	ConsentTemplateID string `json:"consent_template_id"`
	Active            bool   `json:"active"`
}
