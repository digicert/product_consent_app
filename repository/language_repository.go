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

func (lr *LanguageRepository) UnlinkLanguageWithLocale(localeID string, languageID string) (string, error) {
	result, err := lr.DB.Exec("DELETE FROM locale_language WHERE locale_id = ? AND language_id = ?", localeID, languageID)
	if err != nil {
		return "", fmt.Errorf("failed to unlink language with locale: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, language not unlinked")
	}
	return languageID, nil
}

func (lr *LanguageRepository) GetLanguageByID(ID string) (*models.Language, error) {
	language := &models.Language{}
	err := lr.DB.QueryRow("SELECT id, language FROM language WHERE id = ?", ID).Scan(&language.ID, &language.Language)
	if err != nil {
		return nil, fmt.Errorf("failed to get language: %v", err)
	}
	return language, nil
}

func (lr *LanguageRepository) GetAllLanguages(offset, limit int) ([]*models.Language, error) {
	rows, err := lr.DB.Query("SELECT id, language FROM language LIMIT ?, ?", offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get languages: %v", err)
	}
	defer rows.Close()

	var languages []*models.Language
	for rows.Next() {
		var language models.Language
		if err := rows.Scan(&language.ID, &language.Language); err != nil {
			return nil, fmt.Errorf("failed to scan language: %v", err)
		}
		languages = append(languages, &language)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan languages: %v", err)
	}
	return languages, nil
}
