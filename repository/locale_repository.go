package repository

import (
	"database/sql"
	"fmt"

	"github.com/digicert/product-consent-app/models"
	"github.com/google/uuid"
)

type LocaleRepository struct {
	DB *sql.DB
}

func NewLocaleRepository(db *sql.DB) *LocaleRepository {
	return &LocaleRepository{
		DB: db,
	}
}

func (lr *LocaleRepository) UpdateLocale(ID string, locale string) (string, error) {
	result, err := lr.DB.Exec("UPDATE locale SET locale = ? WHERE id = ?", locale, ID)
	if err != nil {
		return "", fmt.Errorf("failed to update locale: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, locale not updated")
	}
	return ID, nil
}

func (lr *LocaleRepository) CreateLocale(locale *models.Locale) (string, error) {
	locale.ID = uuid.New().String()
	_, err := lr.DB.Exec("INSERT INTO locale (id, locale) VALUES (?, ?)", locale.ID, locale.Locale)
	if err != nil {
		return "", fmt.Errorf("Failed to create locale: %v", err)
	}
	return locale.ID, nil
}

func (lr *LocaleRepository) DeleteLocale(ID string) (string, error) {
	result, err := lr.DB.Exec("DELETE FROM locale WHERE id = ?", ID)
	if err != nil {
		return "", fmt.Errorf("failed to delete locale: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no rows affected, locale not deleted")
	}
	return ID, nil
}

func (lr *LocaleRepository) GetLocaleById(ID string) (*models.Locale, error) {
	locale := &models.Locale{}
	err := lr.DB.QueryRow("SELECT id, locale FROM locale WHERE id = ?", ID).Scan(&locale.ID, &locale.Locale)
	if err != nil {
		return nil, fmt.Errorf("failed to get locale: %v", err)
	}
	return locale, nil
}

func (lr *LocaleRepository) GetAllLocales(offset, limit int) ([]*models.Locale, error) {
	rows, err := lr.DB.Query("SELECT id, locale FROM locale LIMIT ?, ?", offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get locales: %v", err)
	}
	defer rows.Close()

	var locales []*models.Locale
	for rows.Next() {
		var locale models.Locale
		if err := rows.Scan(&locale.ID, &locale.Locale); err != nil {
			return nil, fmt.Errorf("failed to scan locale: %v", err)
		}
		locales = append(locales, &locale)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iteration over locales: %v", err)
	}
	return locales, err
}
