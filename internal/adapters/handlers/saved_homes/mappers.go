/* Copyright (C) Fedir Petryk */

package saved_homes

import (
	"strconv"

	"github.com/samber/lo"
	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
)

func mapPaginationRequestToPagination(req paginationRequest) domain.Pagination {
	return domain.Pagination{
		Page: req.Page,
		Size: req.Size,
	}
}

func mapPropertiesToListResponse(props []properties.Property) []*propertyResponse {
	resp := make([]*propertyResponse, 0, len(props))
	for _, prop := range props {
		propMapped := mapPropertyToPropertyResponse(&prop)
		resp = append(resp, propMapped)
	}

	return resp
}

func mapPropertyToPropertyResponse(prop *properties.Property) *propertyResponse {
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

		Appliance:   prop.Appliance,
		Heating:     prop.Heating,
		PetsAllowed: prop.PetsAllowed,

		CreatedAt: strconv.FormatInt(prop.CreatedAt.Unix(), 10),
		UpdatedAt: strconv.FormatInt(prop.UpdatedAt.Unix(), 10),

		IsSavedHome: lo.ToPtr(true),
	}

	if prop.Images != nil {
		propResp.Images = make([]string, 0, len(prop.Images))
		for _, img := range prop.Images {
			propResp.Images = append(propResp.Images, img.ID.String())
		}
	}

	return propResp
}

func propertiesToListResponse(props []properties.Property) []propertyResponse {
	resp := make([]propertyResponse, 0, len(props))
	for _, prop := range props {
		propMapped := propertyToResponse(&prop)
		resp = append(resp, *propMapped)
	}

	return resp
}

func propertyToResponse(prop *properties.Property) *propertyResponse {
	return &propertyResponse{
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
		Description:    prop.Description,
		IsActiveStatus: prop.Active,
		HasImages:      prop.HasImages,
		HasVideo:       prop.HasVideo,
		Has3DTour:      prop.Has3DTour,
		TotalParking:   prop.TotalParking,
		HasAC:          prop.HasAC,

		Images: []string{
			"eda395e8acbe37f588aa97ec17e9e62a-p_e",
			"2c064b023e319850e30a7e0e81f9d24d-p_e",
			"52f11387e2279a98240cee828bf4331b-p_e",
		},
	}
}
