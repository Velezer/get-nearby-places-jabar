package seed

import (
	"jabar-nearby-places/models"
	"jabar-nearby-places/services"
	"jabar-nearby-places/utils"
	"sync"
)

func Begin() {
	seedCategories()
	seedPlaces()
}

func seedCategories() {
	cs := []models.Category{
		{Name: models.CATEGORY_KANTOR_PEM_KABKOTA},
		{Name: models.CATEGORY_RUMAH_SAKIT},
		{Name: models.CATEGORY_SMA},
		{Name: models.CATEGORY_PUSKESMAS},
		{Name: models.CATEGORY_SMP},
		{Name: models.CATEGORY_KANTOR_PEM_KECAMATAN},
		{Name: models.CATEGORY_SD},
		{Name: models.CATEGORY_TEMPAT_IBADAH},
		{Name: models.CATEGORY_KANTOR_PEM_KELURAHANDESA},
	}
	err := services.CategoryService.SaveMany(&cs)
	if err != nil && !utils.ErrDuplicate(err) {
		panic(err)
	}
}

func seedPlaces() {
	cats, err := services.CategoryService.FindAll()
	if err != nil {
		panic(err)
	}
	catmap := map[string]uint{}
	for _, c := range *cats {
		catmap[c.Name] = c.ID
	}

	ws, err := models.LoadWilayah()
	if err != nil {
		panic(err)
	}
	ps := models.GeneratePlaces(ws, catmap)

	var wg sync.WaitGroup
	for i := 0; i < len(ps); i += 5000 {
		j := i + 5000
		if j >= len(ps) {
			j = len(ps)
		}
		data := ps[i:j]
		wg.Add(1)
		go func() {
			err = services.PlaceService.SaveMany(&data)
			if err != nil && !utils.ErrDuplicate(err) {
				panic(err)
			}
			wg.Done()
		}()

	}
	wg.Wait()

}
