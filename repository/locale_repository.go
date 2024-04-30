package repository

import (
	"database/sql"

	"github.com/digicert/product-consent-app/models"
)

type LocaleRepository struct {
	DB *sql.DB
}

func (lr *LocaleRepository) UpdateLocale(locale *models.Locale) error {
	_, err := lr.DB.Exec("UPDATE locale SET locale = ? WHERE id = ?", locale.Locale, locale.ID)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LocaleRepository) CreateLocale(locale *models.Locale) error {
	_, err := lr.DB.Exec("INSERT INTO locale (id, locale) VALUES (?, ?)", locale.ID, locale.Locale)
	if err != nil {
		return err
	}
	return nil
}
