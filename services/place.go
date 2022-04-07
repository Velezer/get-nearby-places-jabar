package services

import (
	"html"
	"jabar-nearby-places/models"
	"jabar-nearby-places/utils"
	"strings"
	"sync"

	"gorm.io/gorm"
)

type place struct {
	db *gorm.DB
}

func (s place) Save(m *models.Place) (*models.Place, error) {
	go CacheService.Clear()
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))

	err := s.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s place) SaveMany(ms *[]models.Place) error {
	go CacheService.Clear()
	return s.db.Create(&ms).Error
}

func (s place) FindAll(categoryId string) (ms *[]models.Place, err error) {
	cache, ok := CacheService.Get("place:findall:" + categoryId).(*[]models.Place)
	if ok {
		return cache, nil
	}
	if categoryId != "" {
		err = s.db.Session(&gorm.Session{PrepareStmt: true}).Find(&ms, models.Place{CategoryID: utils.ParseUint(categoryId, 0)}).Error
	} else {
		err = s.db.Session(&gorm.Session{PrepareStmt: true}).Find(&ms).Error
	}
	CacheService.Set("place:findall:"+categoryId, ms)
	return
}

func (s place) FilterByDistance(ps []models.Place, lat, lon float64) (ms []models.Place) {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, f := range ps {
		wg.Add(1)
		go addIfDistanceLTE5KM(&ms, f, lat, lon, &wg, &m)
	}
	wg.Wait()
	return
}

func addIfDistanceLTE5KM(ms *[]models.Place, p models.Place, lat, lon float64, wg *sync.WaitGroup, m *sync.Mutex) {
	d := utils.DistanceKM(p.Latitude, p.Longitude, lat, lon)
	if d <= 5 {
		m.Lock()
		*ms = append(*ms, p)
		m.Unlock()
	}
	wg.Done()
}
