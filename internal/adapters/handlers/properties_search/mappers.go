/* Copyright (C) Fedir Petryk */

package properties_search

import (
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/serajam/realestate-sample-app/internal/core/domain/search"
)

func mapSearchRequestToSearch(sr *propertySearchRequest) *search.PropertySearchRequest {
	r := &search.PropertySearchRequest{
		City:        sr.City,
		CountryCode: sr.CountryCode,
		BaseFilters: search.BaseFilters{
			Page: sr.Page,
			Size: sr.Size,
			Sort: sr.Sort,
		},
		SearchFilters: search.SearchFilters{
			PriceFrom:     sr.PriceFrom,
			PriceTo:       sr.PriceTo,
			HomeSizeFrom:  sr.HomeSizeFrom,
			HomeSizeTo:    sr.HomeSizeTo,
			LotSizeFrom:   sr.LotSizeFrom,
			LotSizeTo:     sr.LotSizeTo,
			YearBuildFrom: sr.YearBuildFrom,
			YearBuildTo:   sr.YearBuildTo,
			Bathroom:      sr.Bathroom,
			BathroomExact: sr.BathroomExact,
			Bedroom:       sr.Bedroom,
			BedroomExact:  sr.BedroomExact,
			Condition:     sr.Condition,
			HomeType:      sr.HomeType,
			PropertyType:  sr.PropertyType,
			HasAC:         sr.HasAC,
			HasGarage:     sr.HasGarage,
			ParkingNumber: sr.ParkingNumber,
		},
	}

	if sr.Polygon != nil {
		r.Polygon = &search.Polygon{
			TopLat:     sr.Polygon.TopLat,
			TopLong:    sr.Polygon.TopLong,
			BottomLat:  sr.Polygon.BottomLat,
			BottomLong: sr.Polygon.BottomLong,
		}
	}

	return r
}

func propertiesToListResponse(props []properties.Property) []propertyResponse {
	resp := make([]propertyResponse, 0, len(props))
	for _, prop := range props {
		propMapped := propertyToResponse(&prop)
		resp = append(resp, *propMapped)
	}

	return resp
}

func propertiesToMarkerResponse(props []properties.Property) []propertyMarkerResponse {
	resp := make([]propertyMarkerResponse, 0, len(props))
	for _, prop := range props {
		propMapped := propertyToMarkerResponse(&prop)
		resp = append(resp, *propMapped)
	}

	return resp
}

func propertyToResponse(prop *properties.Property) *propertyResponse {
	propResp := &propertyResponse{
		ID:            prop.ID,
		ActualDays:    prop.ActualDays(),
		Location:      location{Lat: prop.Location[0], Long: prop.Location[1]},
		Price:         prop.Price,
		PriceCurrency: prop.PriceCurrency,
		Address: address{
			Country:      prop.Country,
			City:         prop.City,
			State:        prop.State,
			Street:       prop.Street,
			ZipCode:      prop.ZipCode,
			HouseNumber:  prop.HouseNumber,
			Neighborhood: prop.Neighborhood,
		},
		FullAddress:    prop.Address,
		HasGarage:      prop.HasGarage,
		HomeSize:       prop.HomeSize,
		LotSize:        prop.LotSize,
		YearBuild:      prop.YearBuild,
		Bedroom:        prop.Bedroom,
		Bathroom:       prop.Bathroom,
		Floor:          prop.Floor,
		TotalFloors:    prop.TotalFloors,
		PropertyType:   prop.PropertyType,
		HomeType:       prop.HomeType,
		Condition:      prop.Condition,
		BrokerName:     prop.BrokerName,
		IsActiveStatus: prop.Active,
		HasImages:      prop.HasImages,
		HasVideo:       prop.HasVideo,
		Has3DTour:      prop.Has3DTour,
		TotalParking:   prop.TotalParking,
		HasAC:          prop.HasAC,
		Appliance:      prop.Appliance,
		Heating:        prop.Heating,
		PetsAllowed:    prop.PetsAllowed,

		ModelDateTime: common.ModelDateTime{
			CreatedAt: prop.CreatedAt.Unix(),
			UpdatedAt: prop.UpdatedAt.Unix(),
		},
	}

	if prop.Images != nil {
		propResp.Images = make([]string, 0, len(prop.Images))
		for _, img := range prop.Images {
			propResp.Images = append(propResp.Images, img.ID.String())
		}
	}

	return propResp
}

func propertyToMarkerResponse(prop *properties.Property) *propertyMarkerResponse {
	propResp := &propertyMarkerResponse{
		ID:            prop.ID,
		Location:      location{Lat: prop.Location[0], Long: prop.Location[1]},
		Price:         prop.Price,
		PriceCurrency: prop.PriceCurrency,
		Address: address{
			Country:      prop.Country,
			City:         prop.City,
			State:        prop.State,
			Street:       prop.Street,
			ZipCode:      prop.ZipCode,
			HouseNumber:  prop.HouseNumber,
			Neighborhood: prop.Neighborhood,
		},
		Bedroom:  prop.Bedroom,
		Bathroom: prop.Bathroom,
	}

	if prop.Images != nil {
		propResp.Images = make([]string, 0, len(prop.Images))
		for _, img := range prop.Images {
			propResp.Images = append(propResp.Images, img.ID.String())
		}
	}

	// @TODO remove
	if prop.Images == nil {
		propResp.Images = []string{
			"eda395e8acbe37f588aa97ec17e9e62a-p_e",
			"2c064b023e319850e30a7e0e81f9d24d-p_e",
			"52f11387e2279a98240cee828bf4331b-p_e",
		}
	}

	return propResp
}
