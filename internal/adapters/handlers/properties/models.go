/* Copyright (C) Fedir Petryk */

package properties

import (
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
)

type propertyListRequest struct {
	IDs []int `json:"ids" validate:"required,max=10,dive,gte=0,lte=100000000"`
} // @name propertyListRequest

type propertyGetRequest struct {
	ID int `query:"id" validate:"required,gte=0,lte=100000000"`
} // @name propertyGetRequest

type propertyStatusUpdateRequest struct {
	ID     int    `query:"id" validate:"required,gte=0,lte=100000000"`
	Status string `query:"id" validate:"required,gte=0,lte=100000000"`
} // @name propertyGetRequest

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
	LivingSize   float32 `json:"livingSize"`
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
	HasImages      *bool `json:"hasImages"`
	HasGarage      *bool `json:"hasGarage"`
	HasVideo       *bool `json:"hasVideo"`
	Has3DTour      *bool `json:"has3DTour"`
	TotalParking   uint8 `json:"totalParking"`
	HasAC          *bool `json:"hasAC"`

	PetsAllowed *bool `json:"petsAllowed"`
	Appliance   *bool `json:"appliance"`
	Heating     uint8 `json:"heating"`

	Broker broker `json:"broker,omitempty"`

	Images []string `json:"images"`

	common.ModelDateTime
} // @name propertyResponse

type broker struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	CompanyName string `json:"companyName"`
} // @name broker

type address struct {
	Country      string `json:"country" validate:"max=100"`
	City         string `json:"city" validate:"max=100"`
	State        string `json:"state" validate:"max=100"`
	Street       string `json:"street" validate:"max=200"`
	ZipCode      string `json:"zipCode" validate:"max=10"`
	HouseNumber  string `json:"houseNumber" validate:"max=10"`
	Neighborhood string `json:"neighborhood" validate:"max=100"`
} // @name address

type propertyCreateRequest struct {
	Location      location `json:"location" validate:"required"`
	Price         float32  `json:"price" validate:"required,min=0,max=99999999"`
	PriceCurrency string   `json:"priceCurrency" validate:"required,min=0,max=3"`

	FullAddress string  `json:"fullAddress" validate:"required,min=0,max=1000"`
	Address     address `json:"address"`

	LivingSize   float32 `json:"livingSize" validate:"min=0,max=9999"`
	HomeSize     float32 `json:"homeSize" validate:"min=0,max=9999"`
	LotSize      float32 `json:"lotSize" validate:"min=0,max=9999"`
	YearBuild    uint16  `json:"yearBuild" validate:"number,min=0,max=9999"`
	Bedroom      uint8   `json:"bedroom" validate:"number,min=0,max=50"`
	Bathroom     uint8   `json:"bathroom" validate:"number,min=0,max=50"`
	Floor        uint8   `json:"floor" validate:"number,min=0,max=999"`
	TotalFloors  uint8   `json:"totalFloors" validate:"number,min=0,max=999"`
	PropertyType uint8   `json:"propertyType" validate:"required,min=0,max=20"`
	HomeType     uint8   `json:"homeType" validate:"required,min=0,max=20"`
	Condition    uint8   `json:"condition" validate:"required,min=0,max=20"`

	BrokerName  string `json:"brokerName" validate:"required,min=0,max=120"`
	Description string `json:"description" validate:"required,min=0,max=10000"`

	IsActiveStatus *bool `json:"isActiveStatus" validate:"boolean"`
	HasImages      *bool `json:"hasImages" validate:"boolean"`
	HasGarage      *bool `json:"hasGarage" validate:"boolean"`
	HasVideo       *bool `json:"hasVideo" validate:"boolean"`
	Has3DTour      *bool `json:"has3DTour" validate:"boolean"`
	TotalParking   uint8 `json:"totalParking" validate:"min=0,max=100"`
	HasAC          *bool `json:"hasAC" validate:"boolean"`

	PetsAllowed *bool `json:"petsAllowed" validate:"boolean"`
	Appliance   *bool `json:"appliance" validate:"boolean"`
	Heating     uint8 `json:"heating" validate:"min=0,max=3"`
} // @name propertyCreateRequest

type propertyUpdateRequest struct {
	Location      location `json:"location"`
	Price         float32  `json:"price" validate:"min=0,max=99999999"`
	PriceCurrency string   `json:"priceCurrency" validate:"min=0,max=3"`

	FullAddress string  `json:"fullAddress" validate:"min=0,max=1000"`
	Address     address `json:"address"`

	HomeSize     float32 `json:"homeSize" validate:"min=0,max=9999"`
	LotSize      float32 `json:"lotSize" validate:"min=0,max=9999"`
	LivingSize   float32 `json:"livingSize" validate:"min=0,max=9999"`
	YearBuild    uint16  `json:"yearBuild" validate:"number,min=0,max=9999"`
	Bedroom      uint8   `json:"bedroom" validate:"number,min=0,max=50"`
	Bathroom     uint8   `json:"bathroom" validate:"number,min=0,max=50"`
	Floor        uint8   `json:"floor" validate:"number,min=0,max=999"`
	TotalFloors  uint8   `json:"totalFloors" validate:"number,min=1,max=999"`
	PropertyType uint8   `json:"propertyType" validate:"min=0,max=20"`
	HomeType     uint8   `json:"homeType" validate:"min=0,max=20"`
	Condition    uint8   `json:"condition" validate:"min=0,max=20"`

	BrokerName  string `json:"brokerName" validate:"min=0,max=120"`
	Description string `json:"description" validate:"min=0,max=10000"`

	IsActiveStatus *bool `json:"isActiveStatus" validate:"omitempty,boolean"`
	HasImages      *bool `json:"hasImages" validate:"omitempty,boolean"`
	HasGarage      *bool `json:"hasGarage" validate:"omitempty,boolean"`
	HasVideo       *bool `json:"hasVideo" validate:"omitempty,boolean"`
	Has3DTour      *bool `json:"has3DTour" validate:"omitempty,boolean"`
	TotalParking   uint8 `json:"totalParking" validate:"min=0,max=100"`
	HasAC          *bool `json:"hasAC" validate:"omitempty,boolean"`

	PetsAllowed *bool `json:"petsAllowed" validate:"omitempty,boolean"`
	Appliance   *bool `json:"appliance" validate:"omitempty,boolean"`
	Heating     uint8 `json:"heating" validate:"min=0,max=3"`
} // @name propertyCreateRequest

type location struct {
	Lat  float64 `json:"lat" validate:"required,latitude"`
	Long float64 `json:"long" validate:"required,longitude"`
} // @name propertyLocation
