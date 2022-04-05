package models

import (
	"jabar-nearby-places/dataset"
	"jabar-nearby-places/utils"
)

type Wilayah struct {
	Name      string `csv:""`
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

	for _, v := range jabars {
		ws = append(ws, Wilayah{
			Name:      v.Bps_kota_nama,
			Level:     "Kabupaten/Kota",
			Code:      v.Kemendagri_kota_kode,
			Latitude:  utils.ParseFloat64(v.Latitude, -99),
			Longitude: utils.ParseFloat64(v.Longitude, -99),
		})
		ws = append(ws, Wilayah{
			Name:      v.Kemendagri_kecamatan_nama,
			Level:     "Kecamatan",
			Code:      v.Kemendagri_kecamatan_kode,
			Latitude:  utils.ParseFloat64(v.Latitude, -99),
			Longitude: utils.ParseFloat64(v.Longitude, -99),
		})
		ws = append(ws, Wilayah{
			Name:      v.Kemendagri_kelurahan_nama,
			Level:     "Kelurahan/Desa",
			Code:      v.Kemendagri_kelurahan_kode,
			Latitude:  utils.ParseFloat64(v.Latitude, -99),
			Longitude: utils.ParseFloat64(v.Longitude, -99),
		})
		ws = unique(ws)
	}
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
