package services

import (
	"html"
	"jabar-nearby-places/models"
	"jabar-nearby-places/utils"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
)

type place struct {
	db *gorm.DB
}

func (s place) Save(m *models.Place) (*models.Place, error) {
	go CacheService.ClearAll()
	m.Name = html.EscapeString(strings.TrimSpace(m.Name))

	err := s.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}
func (s place) SaveMany(ms *[]models.Place) error {
	go CacheService.ClearAll()
	return s.db.Create(&ms).Error
}

func (s place) FindAll(categoryId uint) (ms *[]models.Place, err error) {
	cid := strconv.FormatUint(uint64(categoryId), 10)
	cache, ok := CacheService.Get("place:findall:" + cid).(*[]models.Place)
	if ok {
		return cache, nil
	}
	err = s.db.Session(&gorm.Session{PrepareStmt: true}).Find(&ms, models.Place{CategoryID: categoryId}).Error
	CacheService.Set("place:findall:"+cid, ms, 60*5) // cache the result for 5 minutes
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
