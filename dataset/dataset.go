package dataset

import (
	"io/ioutil"
	"os"

	"github.com/jszwec/csvutil"
)

type Jabar struct {
	CityCode     string `csv:"kemendagri_kota_kode"`
	DistrictCode string `csv:"kemendagri_kecamatan_kode"`
	VillageCode  string `csv:"kemendagri_kelurahan_kode"`

	CityName     string `csv:"bps_kota_nama"`
	DistrictName string `csv:"kemendagri_kecamatan_nama"`
	VillageName  string `csv:"kemendagri_kelurahan_nama"`

	Latitude  string `csv:"latitude"`
	Longitude string `csv:"longitude"`
}

func Load() (jabars []Jabar, err error) {
	f, err := os.Open("dataset/diskominfo-od_kode_wilayah_dan_nama_wilayah_desa_kelurahan_data.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = csvutil.Unmarshal(d, &jabars)
	if err != nil {
		return nil, err
	}
	return jabars, nil
}
