package main

import "jabar-nearby-places/routes"

func main() {
	// seed.Begin()

	r := routes.SetupRouter()
	r.Run()
}
