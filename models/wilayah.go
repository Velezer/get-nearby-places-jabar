package models

import (
	"jabar-nearby-places/dataset"
	"jabar-nearby-places/utils"
	"strings"
	"sync"
)

const (
	LEVEL_KABKOTA       = "Kabupaten/Kota"
	LEVEL_KECAMATAN     = "Kecamatan"
	LEVEL_KELURAHANDESA = "Kelurahan/Desa"
)

// this model is not included in database
type Wilayah struct {
	Name      string
	Level     string
	Code      string
	Latitude  float64
	Longitude float64

	CityName     string
	DistrictName string
}

func LoadWilayah() (ws []Wilayah, err error) {
	jabars, err := dataset.Load()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var m sync.Mutex
	for _, v := range jabars {
		wg.Add(3)
		go addWs(&ws, Wilayah{
			Name:         v.CityName,
			Level:        LEVEL_KABKOTA,
			Code:         v.CityCode,
			Latitude:     utils.ParseFloat64(v.Latitude, -99),
			Longitude:    utils.ParseFloat64(v.Longitude, -99),
			CityName:     v.CityName,
			DistrictName: "Kecamatan " + v.DistrictName,
		}, &wg, &m)
		go addWs(&ws, Wilayah{
			Name:         v.DistrictName,
			Level:        LEVEL_KECAMATAN,
			Code:         v.DistrictCode,
			Latitude:     utils.ParseFloat64(v.Latitude, -99),
			Longitude:    utils.ParseFloat64(v.Longitude, -99),
			CityName:     v.CityName,
			DistrictName: "Kecamatan " + v.DistrictName,
		}, &wg, &m)
		go addWs(&ws, Wilayah{
			Name:         v.VillageName,
			Level:        LEVEL_KELURAHANDESA,
			Code:         v.VillageCode,
			Latitude:     utils.ParseFloat64(v.Latitude, -99),
			Longitude:    utils.ParseFloat64(v.Longitude, -99),
			CityName:     v.CityName,
			DistrictName: "Kecamatan " + v.DistrictName,
		}, &wg, &m)
	}
	wg.Wait()

	ws = uniqueWs(ws)
	return
}

func uniqueWs(slice []Wilayah) []Wilayah {
	keys := make(map[string]bool)
	list := []Wilayah{}
	for _, entry := range slice {
		if _, value := keys[entry.Code]; !value {
			keys[entry.Code] = true
			list = append(list, entry)
		}
	}
	return list
}

func addWs(ws *[]Wilayah, w Wilayah, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	if len(w.Name) > 0 && !strings.Contains(w.Name, "BELUM TERIDENTIFIKASI") {
		*ws = append(*ws, w)
	}
	m.Unlock()
	wg.Done()
}
