package dataset

import (
	"io/ioutil"
	"os"

	"github.com/jszwec/csvutil"
)

type Jabar struct {
	// id
	// kemendagri_provinsi_kode
	Kemendagri_kota_kode      string `csv:"kemendagri_kota_kode"`
	Kemendagri_kecamatan_kode string `csv:"kemendagri_kecamatan_kode"`
	Kemendagri_kelurahan_kode string `csv:"kemendagri_kelurahan_kode"`
	// Kemendagri_provinsi_nama  string `csv:"kemendagri_provinsi_nama"`
	// Kemendagri_kota_nama      string `csv:"kemendagri_kota_nama"`
	Kemendagri_kecamatan_nama string `csv:"kemendagri_kecamatan_nama"`
	Kemendagri_kelurahan_nama string `csv:"kemendagri_kelurahan_nama"`
	// bps_provinsi_kode
	// bps_kota_kode
	// bps_kecamatan_kode
	// bps_kelurahan_kode
	// Bps_provinsi_nama  string `csv:"bps_provinsi_nama"`
	Bps_kota_nama string `csv:"bps_kota_nama"`
	// Bps_kecamatan_nama string `csv:"bps_kecamatan_nama"`
	// Bps_kelurahan_nama string `csv:"bps_kelurahan_nama"`
	Latitude  string `csv:"latitude"`
	Longitude string `csv:"longitude"`
	// kode_pos
	// Status_adm string `csv:"status_adm"`
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
