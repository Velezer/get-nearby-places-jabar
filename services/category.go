package services

import (
	"jabar-nearby-places/models"

	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

func (s category) FindAll() (categories *[]models.Category, err error) {
	err = s.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return
}
func (s category) FindOne(name string) (category *models.Category, err error) {
	err = s.db.First(&category, models.Category{Name: name}).Error
	if err != nil {
		return nil, err
	}

	return
}
