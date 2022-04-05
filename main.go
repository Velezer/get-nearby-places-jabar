package main

import (
	"jabar-nearby-places/routes"
)

func main() {

	r := routes.SetupRouter()
	r.Run()
}
