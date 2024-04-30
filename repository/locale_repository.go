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

func (lr *LocaleRepository) UpdateLocale(locale *models.Locale) error {
	_, err := lr.DB.Exec("UPDATE locale SET locale = ? WHERE id = ?", locale.Locale, locale.ID)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LocaleRepository) CreateLocale(locale *models.Locale) (string, error) {
	locale.ID = uuid.New().String()
	_, err := lr.DB.Exec("INSERT INTO locale (id, locale) VALUES (?, ?)", locale.ID, locale.Locale)
	if err != nil {
		return "", fmt.Errorf("Failed to create locale: %v", err)
	}
	return locale.ID, nil
}
