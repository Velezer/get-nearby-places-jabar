package main

import (
	"jabar-nearby-places/routes"
	"jabar-nearby-places/seed"
	"jabar-nearby-places/services"
)

func main() {
	go seed.Begin()                      // seed if never been seeded
	go services.PlaceService.FindAll("") // will cache findall result for increasing search performance

	r := routes.SetupRouter()
	r.Run()
}
