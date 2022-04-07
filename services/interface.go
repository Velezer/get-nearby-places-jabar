package services

import "jabar-nearby-places/models"

type (
	placeIface interface {
		Save(m *models.Place) (*models.Place, error)
		SaveMany(m *[]models.Place) error
		FindAll(categoryId uint) (ps *[]models.Place, err error)
		FilterByDistance(ps []models.Place, lat, lon float64) (ms []models.Place)
	}

	categoryIface interface {
		FindAll() (categories *[]models.Category, err error)
		FindOne(name string) (categories *models.Category, err error)
		SaveMany(ms *[]models.Category) error
	}

	cacheIface interface {
		Set(key string, value interface{}, ttls uint)
		Get(key string) interface{}
		ClearAll()
		Clear(key string)
	}
)
