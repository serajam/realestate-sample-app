/* Copyright (C) Fedir Petryk */

package search_filters

type searchFiltersRequest struct {
	Name           string              `json:"name" validate:"omitempty,max=300"`
	EmailFrequency uint8               `json:"emailFrequency" validate:"omitempty,number,min=1,max=10"`
	Sort           uint8               `json:"sort" validate:"omitempty,number,min=1,max=10"`
	Subscribed     *bool               `json:"subscribed" validate:"omitempty,boolean"`
	Filters        *savedSearchFilters `json:"filterValues"`
	Polygon        *polygon            `json:"coordinatesRect"`
} // @name searchFiltersRequest

type searchFiltersResponse struct {
	ID             int                `json:"id"`
	Name           string             `json:"name" validate:"max=300"`
	EmailFrequency uint8              `json:"emailFrequency"`
	Sort           uint8              `json:"sort,omitempty"`
	Subscribed     bool               `json:"subscribed" validate:"omitempty,boolean"`
	Filters        savedSearchFilters `json:"filterValues,omitempty"`
	Polygon        *polygon           `json:"coordinatesRect,omitempty"`
} // @name searchFiltersResponse

type savedSearchFilters struct {
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

	HasAC          *bool  `json:"hasAC,omitempty" validate:"omitempty,boolean"`
	MustHaveGarage bool   `json:"mustHaveGarage,omitempty" validate:"omitempty,boolean"`
	ParkingNumber  *uint8 `json:"parkingNumber,omitempty" validate:"omitempty,number"`
} // @name savedSearchFilters

type searchFiltersCreateResponse struct {
	ID int `json:"id"`
} // @name searchFiltersCreateResponse

type polygon struct {
	TopLat     string `json:"topLat" validate:"required_without=City,omitempty,latitude"`
	TopLong    string `json:"topLong" validate:"required_without=City,omitempty,latitude"`
	BottomLat  string `json:"botLat" validate:"required_without=City,omitempty,latitude"`
	BottomLong string `json:"botLong" validate:"required_without=City,omitempty,latitude"`
} // @name polygon
