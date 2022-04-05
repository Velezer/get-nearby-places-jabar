package main

import (
	"fmt"
	"jabar-nearby-places/models"
	"log"
)

func main() {
	wilayahs, err := models.LoadWilayah()
	if err != nil {
		panic(err)
	}
	for _, v := range wilayahs {
		fmt.Println(v)
	}
	log.Fatal()

	// r := routes.SetupRouter()
	// r.Run()
}
