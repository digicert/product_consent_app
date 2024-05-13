package models

type ConsentTemplate struct {
	ID               string `json:"id"`
	LocaleLanguageID string `json:"locale_language_id"`
	TemplatePDF      []byte `json:"template_pdf"`
}
