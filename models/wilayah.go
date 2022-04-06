package models

import (
	"jabar-nearby-places/dataset"
	"jabar-nearby-places/utils"
	"sync"
)

type Wilayah struct {
	Name      string
	Level     string
	Code      string
	Latitude  float64
	Longitude float64
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
			Name:      v.Bps_kota_nama,
			Level:     "Kabupaten/Kota",
			Code:      v.Kemendagri_kota_kode,
			Latitude:  utils.ParseFloat64(v.Latitude, -99),
			Longitude: utils.ParseFloat64(v.Longitude, -99),
		}, &wg, &m)
		go addWs(&ws, Wilayah{
			Name:      v.Kemendagri_kecamatan_nama,
			Level:     "Kecamatan",
			Code:      v.Kemendagri_kecamatan_kode,
			Latitude:  utils.ParseFloat64(v.Latitude, -99),
			Longitude: utils.ParseFloat64(v.Longitude, -99),
		}, &wg, &m)
		go addWs(&ws, Wilayah{
			Name:      v.Kemendagri_kelurahan_nama,
			Level:     "Kelurahan/Desa",
			Code:      v.Kemendagri_kelurahan_kode,
			Latitude:  utils.ParseFloat64(v.Latitude, -99),
			Longitude: utils.ParseFloat64(v.Longitude, -99),
		}, &wg, &m)
	}
	wg.Wait()

	ws = unique(ws)
	return
}

func unique(slice []Wilayah) []Wilayah {
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
	*ws = append(*ws, w)
	m.Unlock()
	wg.Done()
}
