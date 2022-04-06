package main

import (
	"fmt"
	"jabar-nearby-places/models"
)

func main() {
	ws, err := models.LoadWilayah()
	if err != nil {
		panic(err)
	}

	ps := models.GeneratePlaces(ws)
	for _, v := range ps {
		fmt.Println(v.Name)
	}
	// r := routes.SetupRouter()
	// r.Run()
}
