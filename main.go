package main

import (
	"jabar-nearby-places/routes"
	"jabar-nearby-places/seed"
	"jabar-nearby-places/services"
)

func main() {
	go func() { // seed if never been seeded
		rows, err := services.PlaceService.CountRows()
		if err != nil {
			panic(err)
		}
		if rows < 104668 { // check if rows has the amount of data
			seed.Begin()
		}
	}()

	r := routes.SetupRouter()
	r.Run()
}
