package services

import (
	"html"
	"jabar-nearby-places/models"
	"strings"

	"gorm.io/gorm"
)

type place struct {
	db *gorm.DB
}

func (s place) Save(m *models.Place) (*models.Place, error) {
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))

	err := s.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s place) SaveMany(ms *[]models.Place) error {
	return s.db.Create(&ms).Error
}
