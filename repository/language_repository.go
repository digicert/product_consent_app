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

func (lr *LanguageRepository) UpdateLanguage(ID string, language string) (string, error) {
	result, err := lr.DB.Exec("UPDATE language SET language = ? WHERE id = ?", language, ID)
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
	return ID, nil
}

func (lr *LanguageRepository) DeleteLanguage(ID string) (string, error) {
	result, err := lr.DB.Exec("DELETE FROM language WHERE id = ?", ID)
	if err != nil {
		return "", fmt.Errorf("failed to delete language: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, language not deleted")
	}
	return ID, nil
}

func (lr *LanguageRepository) CreateLanguage(language *models.Language) (string, error) {
	language.ID = uuid.New().String()
	_, err := lr.DB.Exec("INSERT INTO language (id, language) VALUES (?, ?)", language.ID, language.Language)
	if err != nil {
		return "", fmt.Errorf("failed to create language: %v", err)
	}
	return language.ID, nil
}

func (lr *LanguageRepository) LinkLanguageWithLocale(ID string, localeID string, languageID string) (string, error) {
	ID = uuid.New().String()
	_, err := lr.DB.Exec("INSERT INTO locale_language (id, locale_id, language_id) VALUES (?, ?, ?)", ID, localeID, languageID)
	if err != nil {
		return "", fmt.Errorf("failed to link language with locale: %v", err)
	}
	return languageID, nil
}
