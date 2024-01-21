/* Copyright (C) Fedir Petryk */

package saved_homes

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

	BrokerName  string `json:"brokerName"`
	Description string `json:"description"`

	IsActiveStatus *bool `json:"isActiveStatus"`
	IsSavedHome    *bool `json:"isSavedHome"`
	HasImages      *bool `json:"hasImages"`
	HasGarage      *bool `json:"hasGarage"`
	HasVideo       *bool `json:"hasVideo"`
	Has3DTour      *bool `json:"has3DTour"`
	TotalParking   uint8 `json:"totalParking"`
	HasAC          *bool `json:"hasAC"`

	PetsAllowed *bool `json:"petsAllowed"`
	Appliance   *bool `json:"appliance"`
	Heating     uint8 `json:"heating"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Images []string `json:"images"`
} // @name propertyResponse

type address struct {
	Country      string `json:"country"`
	City         string `json:"city"`
	State        string `json:"state"`
	Street       string `json:"street"`
	ZipCode      string `json:"zipCode"`
	HouseNumber  string `json:"houseNumber"`
	Neighborhood string `json:"neighborhood"`
} // @name address

type location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
} // @name propertyLocationHome

type paginationRequest struct {
	Page int `json:"page" validate:"number,max=100000" json:"page,omitempty"`
	Size int `json:"size" validate:"number,max=1000" json:"size,omitempty"`
} // @name paginationRequest
