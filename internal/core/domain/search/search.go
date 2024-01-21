/* Copyright (C) Fedir Petryk */

package search

type BaseFilters struct {
	Page int
	Size int
	Sort int
}

type PropertySearchRequest struct {
	City        string
	CountryCode string

	Polygon *Polygon

	BaseFilters
	SearchFilters
}

type SearchFilters struct {
	PriceFrom     float32
	PriceTo       float32
	HomeSizeFrom  float32
	HomeSizeTo    float32
	LotSizeFrom   float32
	LotSizeTo     float32
	YearBuildFrom uint16
	YearBuildTo   uint16

	Bathroom      uint8
	BathroomExact bool

	Bedroom      uint8
	BedroomExact bool

	Condition    []int
	HomeType     []int
	PropertyType uint8

	HasAC         *bool
	HasGarage     bool
	ParkingNumber *uint8
}

type Polygon struct {
	TopLat     string
	TopLong    string
	BottomLat  string
	BottomLong string
}
