package models

import (
	"jabar-nearby-places/mlocal"
	"sync"
)

type Place struct {
	BaseModel
	Name       string   `json:"name" gorm:"unique"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"-"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
}

func addPlaces(ps *[]Place, p Place, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*ps = append(*ps, p)
	m.Unlock()
	wg.Done()
}

func PreparePlaces(pslocal []mlocal.Place, catmap map[string]uint) (ps []Place) {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, p := range pslocal {
		wg.Add(1)
		addPlaces(&ps, Place{
			Name:       p.Name,
			CategoryID: catmap[p.Category],
			Latitude:   p.Latitude,
			Longitude:  p.Longitude,
		}, &wg, &m)
	}
	wg.Wait()
	return
}
