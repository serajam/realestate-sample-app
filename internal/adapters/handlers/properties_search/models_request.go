/* Copyright (C) Fedir Petryk */

package properties_search

import (
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

type propertyGetRequest struct {
	ID int `query:"id" validate:"required,gte=0,lte=100000000"`
} // @name propertyGetRequest

type propertySearchRequest struct {
	City        string `json:"city" validate:"omitempty,max=255,required_without=polygon" json:"city,omitempty"`
	CountryCode string `json:"countryCode" validate:"omitempty,max=2,required_without_all=polygon" json:"countryCode,omitempty"`

	Polygon *polygon `json:"polygon" validate:"required_without=City"`

	common.BaseSearchRequest
	searchFilters
} // @name propertySearchRequest

type searchFilters struct {
	PriceFrom     float32 `json:"priceFrom,omitempty" validate:"omitempty,number"`
	PriceTo       float32 `json:"priceTo,omitempty" validate:"omitempty,number"`
	HomeSizeFrom  float32 `json:"homeSizeFrom,omitempty" validate:"omitempty,number"`
	HomeSizeTo    float32 `json:"homeSizeTo,omitempty" validate:"omitempty,number"`
	LotSizeFrom   float32 `json:"lotSizeFrom,omitempty" validate:"omitempty,number"`
	LotSizeTo     float32 `json:"lotSizeTo,omitempty" validate:"omitempty,number"`
	YearBuildFrom uint16  `json:"yearBuildFrom,omitempty" validate:"omitempty,number"`
	YearBuildTo   uint16  `json:"yearBuildTo,omitempty" validate:"omitempty,number"`

	Bathroom      uint8 `json:"bathroom,omitempty" validate:"omitempty,number"`
	BathroomExact bool  `json:"bathroomExact,omitempty" validate:"omitempty,boolean"`

	Bedroom      uint8 `json:"bedroom,omitempty" validate:"omitempty,number"`
	BedroomExact bool  `json:"bedroomExact,omitempty" validate:"omitempty,boolean"`

	Condition    []int `json:"condition,omitempty" validate:"omitempty,dive,number"`
	HomeType     []int `json:"homeType,omitempty" validate:"omitempty,dive,number"`
	PropertyType uint8 `json:"propertyType,omitempty" validate:"omitempty,number"`

	HasAC         *bool  `json:"hasAC,omitempty" validate:"omitempty,boolean"`
	HasGarage     bool   `json:"hasGarage,omitempty" validate:"omitempty,boolean"`
	ParkingNumber *uint8 `json:"parkingNumber,omitempty" validate:"omitempty,number"`
}

type polygon struct {
	TopLat     string `json:"topLat" validate:"required_without=City,omitempty,latitude"`
	TopLong    string `json:"topLong" validate:"required_without=City,omitempty,latitude"`
	BottomLat  string `json:"botLat" validate:"required_without=City,omitempty,latitude"`
	BottomLong string `json:"botLong" validate:"required_without=City,omitempty,latitude"`
} // @name polygon
