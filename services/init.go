package services

import (
	"jabar-nearby-places/database"

	"gorm.io/gorm"
)

var db *gorm.DB
var PlaceService placeIface
var CategoryService categoryIface

func init() {
	if db == nil {
		db = database.ConnectDatabase()
	}

	if CategoryService == nil {
		CategoryService = &category{db}
	}

	if PlaceService == nil {
		PlaceService = &place{db}
	}

}
