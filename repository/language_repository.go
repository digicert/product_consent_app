package repository

import (
	"database/sql"
	"fmt"

	"github.com/digicert/product-consent-app/models"
	"github.com/google/uuid"
)

type LanguageRepository struct {
	DB *sql.DB
}

func NewLanguageRepository(db *sql.DB) *LanguageRepository {
	return &LanguageRepository{
		DB: db,
	}
}

func (lr *LanguageRepository) UpdateLanguage(language *models.Language) (string, error) {
	language.ID = uuid.New().String()
	result, err := lr.DB.Exec("UPDATE language SET language = ? WHERE id = ?", language.Language, language.ID)
	if err != nil {
		return "", fmt.Errorf("failed to update language: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, language not updated")
	}
	return language.ID, nil
}

func (lr *LanguageRepository) CreateLanguage(language *models.Language) (string, error) {
	language.ID = uuid.New().String()
	_, err := lr.DB.Exec("INSERT INTO language (id, language) VALUES (?, ?)", language.ID, language.Language)
	if err != nil {
		return "", fmt.Errorf("Failed to create language: %v", err)
	}
	return language.ID, nil
}

func (lr *LanguageRepository) LinkLanguageWithLocale(languageID, localeID string) (string, error) {
	_, err := lr.DB.Exec("INSERT INTO language_locale (language_id, locale_id) VALUES (?, ?)", languageID, localeID)
	if err != nil {
		return "", fmt.Errorf("Failed to link language with locale: %v", err)
	}
	return languageID, nil
}
