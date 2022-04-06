package main

import (
	"fmt"
	"jabar-nearby-places/models"
)

func main() {
	wilayahs, err := models.LoadWilayah()
	if err != nil {
		panic(err)
	}
	for _, v := range wilayahs {
		fmt.Println(v)
	}

	// r := routes.SetupRouter()
	// r.Run()
}
