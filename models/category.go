package models

const (
	CATEGORY_KANTOR_PEM_KABKOTA       = "Kantor Pemerintah Kabupaten/Kota"
	CATEGORY_RUMAH_SAKIT              = "Rumah Sakit"
	CATEGORY_SMA                      = "Sekolah Menengah Atas"
	CATEGORY_PUSKESMAS                = "Puskesmas"
	CATEGORY_SMP                      = "Sekolah Menengah Pertama"
	CATEGORY_KANTOR_PEM_KECAMATAN     = "Kantor Pemerintah Kecamatan"
	CATEGORY_SD                       = "Sekolah Dasar"
	CATEGORY_TEMPAT_IBADAH            = "Tempat Ibadah"
	CATEGORY_KANTOR_PEM_KELURAHANDESA = "Kantor Pemerintah Kelurahan/Desa"
)

type Category struct {
	BaseModel

	Name string `json:"name" gorm:"unique"`
}
