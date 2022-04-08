package main

import (
	"jabar-nearby-places/routes"
	"jabar-nearby-places/seed"
)

func main() {
	go seed.Begin() // seed if never been seeded

	r := routes.SetupRouter()
	r.Run()
}
