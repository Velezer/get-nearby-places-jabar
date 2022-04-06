package main

import (
	"fmt"
	"jabar-nearby-places/mlocal"
	"jabar-nearby-places/models"
	"jabar-nearby-places/services"
	"sync"
)

func main() {
	seedPlaces()
	// r := routes.SetupRouter()
	// r.Run()
}

func seedPlaces() {
	wslocal, _ := mlocal.LoadWilayah()
	pslocal := mlocal.GeneratePlaces(wslocal)
	cats, err := services.CategoryService.FindAll()
	if err != nil {
		panic(err)
	}
	catmap := map[string]uint{}
	for _, c := range *cats {
		catmap[c.Name] = c.ID
	}

	ps := models.PreparePlaces(pslocal, catmap)
	fmt.Println(len(wslocal))
	fmt.Println(len(pslocal))
	fmt.Println(len(ps))

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
