package models

type Place struct {
	baseModel
	Name     string
	Category Category
	Latitude  float64
	Longitude float64
}
