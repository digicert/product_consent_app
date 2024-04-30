package models

type ConsentTemplate struct {
	ID         string `json:"id"`
	LocaleID   string `json:"locale_id"`
	LanguageID string `json:"language_id"`
}
