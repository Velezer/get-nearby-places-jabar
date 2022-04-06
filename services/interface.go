package services

import "jabar-nearby-places/models"

type placeIface interface {
	Save(m *models.Place) (*models.Place, error)
	SaveMany(m *[]models.Place) error
	FindAll() (ps *[]models.Place, err error)
	FilterByDistance(ps []models.Place, lat, lon float64) (ms []models.Place)
}

type categoryIface interface {
	FindAll() (categories *[]models.Category, err error)
	FindOne(name string) (categories *models.Category, err error)
	SaveMany(ms *[]models.Category) error
}
