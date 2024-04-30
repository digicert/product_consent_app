package models

type ClientConsent struct {
	ID                string `json:"id"`
	ProductTemplateID string `json:"product_template_id"`
	IndividualID      string `json:"individual_id"`
	Date              string `json:"date"`
	OptoutReason      string `json:"optout_reason"`
}
