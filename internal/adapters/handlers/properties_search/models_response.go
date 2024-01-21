/* Copyright (C) Fedir Petryk */

package properties_search

import "github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"

type location struct {
	Lat  float64 `json:"lat" validate:"required,latitude"`
	Long float64 `json:"long" validate:"required,longitude"`
} // @name propertyLocationSearch

type propertyResponse struct {
	ID         int `json:"id"`
	ActualDays int `json:"actualDays"`

	Location      location `json:"location"`
	Price         float32  `json:"price"`
	PriceCurrency string   `json:"priceCurrency"`

	FullAddress string  `json:"fullAddress"`
	Address     address `json:"address"`

	HomeSize     float32 `json:"homeSize"`
	LotSize      float32 `json:"lotSize"`
	YearBuild    uint16  `json:"yearBuild"`
	Bedroom      uint8   `json:"bedroom"`
	Bathroom     uint8   `json:"bathroom"`
	Floor        uint8   `json:"floor"`
	TotalFloors  uint8   `json:"totalFloors"`
	PropertyType uint8   `json:"propertyType"`
	HomeType     uint8   `json:"homeType"`
	Condition    uint8   `json:"condition"`

	BrokerName string `json:"brokerName"`

	IsActiveStatus *bool `json:"isActiveStatus"`
	HasImages      *bool `json:"hasImages"`
	HasGarage      *bool `json:"hasGarage"`
	HasVideo       *bool `json:"hasVideo"`
	Has3DTour      *bool `json:"has3DTour"`
	TotalParking   uint8 `json:"totalParking"`
	HasAC          *bool `json:"hasAC"`

	PetsAllowed *bool `json:"petsAllowed"`
	Appliance   *bool `json:"appliance"`
	Heating     uint8 `json:"heating"`

	common.ModelDateTime

	Images []string `json:"images"`
} // @name propertyResponse

type propertyMarkerResponse struct {
	ID            int      `json:"id"`
	Location      location `json:"location"`
	Price         float32  `json:"price"`
	PriceCurrency string   `json:"priceCurrency"`
	Address       address  `json:"address"`
	Bedroom       uint8    `json:"bedroom"`
	Bathroom      uint8    `json:"bathroom"`

	Images []string `json:"images"`
} // @name propertyMarkerResponse

type address struct {
	Country      string `json:"country" validate:"max=100"`
	City         string `json:"city" validate:"max=100"`
	State        string `json:"state" validate:"max=100"`
	Street       string `json:"street" validate:"max=200"`
	ZipCode      string `json:"zipCode" validate:"max=10"`
	HouseNumber  string `json:"houseNumber" validate:"max=10"`
	Neighborhood string `json:"neighborhood" validate:"max=100"`
} // @name address
