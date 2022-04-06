package seed

import (
	"jabar-nearby-places/models"
	"jabar-nearby-places/services"
	"sync"
)

func seedPlaces() {
	cats, err := services.CategoryService.FindAll()
	if err != nil {
		panic(err)
	}
	catmap := map[string]uint{}
	for _, c := range *cats {
		catmap[c.Name] = c.ID
	}

	ws, _ := models.LoadWilayah()
	ps := models.GeneratePlaces(ws, catmap)

	var wg sync.WaitGroup
	for i := 0; i < len(ps); i += 9000 {
		j := i + 9000
		if j >= len(ps) {
			j = len(ps)
		}
		data := ps[i:j]
		wg.Add(1)
		go func() {
			err = services.PlaceService.SaveMany(&data)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()

	}
	wg.Wait()

}
