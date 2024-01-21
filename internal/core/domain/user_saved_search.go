/* Copyright (C) Fedir Petryk */

package domain

import (
	"time"
)

type UserSearchFilters struct {
	ID             int       `bun:"type:id,pk,autoincrement"`
	UserID         int       `bun:"user_id"`
	Name           string    `bun:"name"`
	Sort           uint8     `bun:"sort"`
	Subscribed     *bool     `bun:"subscribed"`
	EmailFrequency uint8     `bun:"email_frequency"`
	CreatedAt      time.Time `bun:",nullzero,default:now()"`
	UpdatedAt      time.Time `bun:",default:current_timestamp"`

	Filters *SavedSearchesFilters `bun:"filters,type:jsonb"`
	Polygon *SearchPolygon        `bun:"polygon,type:jsonb"`
}

type SearchPolygon struct {
	TopLat     string `json:"topLat,omitempty"`
	TopLong    string `json:"topLong,omitempty"`
	BottomLat  string `json:"bottomLat,omitempty"`
	BottomLong string `json:"bottomLong,omitempty"`
}

type SavedSearchesFilters struct {
	PriceFrom     float32 `json:"priceFrom,omitempty"`
	PriceTo       float32 `json:"priceTo,omitempty"`
	HomeSizeFrom  float32 `json:"homeSizeFrom,omitempty"`
	HomeSizeTo    float32 `json:"homeSizeTo,omitempty"`
	LotSizeFrom   float32 `json:"lotSizeFrom,omitempty"`
	LotSizeTo     float32 `json:"lotSizeTo,omitempty"`
	YearBuildFrom uint16  `json:"yearBuildFrom,omitempty"`
	YearBuildTo   uint16  `json:"yearBuildTo,omitempty"`

	Bathroom      uint8 `json:"bathroom,omitempty"`
	BathroomExact bool  `json:"bathroomExact,omitempty"`

	Bedroom      uint8 `json:"bedroom,omitempty"`
	BedroomExact bool  `json:"bedroomExact,omitempty"`

	Condition    []int `json:"condition,omitempty" bun:"polygon,type:jsonb,json_use_number"`
	HomeType     []int `json:"homeType,omitempty" bun:"polygon,type:jsonb,json_use_number"`
	PropertyType uint8 `json:"propertyType,omitempty"`

	HasAC          *bool  `json:"hasAC,omitempty"`
	MustHaveGarage bool   `json:"mustHaveGarage,omitempty"`
	ParkingNumber  *uint8 `json:"parkingNumber,omitempty"`
}
