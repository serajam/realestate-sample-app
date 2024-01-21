/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/serajam/realestate-sample-app/internal/adapters/handlers/common"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/serajam/realestate-sample-app/internal/core/domain/search"
)

func propertiesToListResponse(props []properties.Property) []propertyResponse {
	resp := make([]propertyResponse, 0, len(props))
	for _, prop := range props {
		propMapped := propertyToResponse(&prop)
		resp = append(resp, *propMapped)
	}

	return resp
}

func createPropertyToDomain(ctx echo.Context, prop *propertyCreateRequest) (*properties.Property, error) {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return nil, errors.New("invalid user id")
	}

	return &properties.Property{
		UserID:        userID,
		Location:      properties.Point{prop.Location.Lat, prop.Location.Long},
		Price:         prop.Price,
		PriceCurrency: prop.PriceCurrency,
		Address:       prop.FullAddress,
		Country:       prop.Address.Country,
		City:          prop.Address.City,
		State:         prop.Address.State,
		Street:        prop.Address.Street,
		ZipCode:       prop.Address.ZipCode,
		HouseNumber:   prop.Address.HouseNumber,
		Neighborhood:  prop.Address.Neighborhood,
		HasGarage:     prop.HasGarage,
		LivingSize:    prop.LivingSize,
		LotSize:       prop.LotSize,
		HomeSize:      prop.HomeSize,
		YearBuild:     prop.YearBuild,
		Bedroom:       prop.Bedroom,
		Bathroom:      prop.Bathroom,
		Floor:         prop.Floor,
		TotalFloors:   prop.TotalFloors,
		PropertyType:  prop.PropertyType,
		HomeType:      prop.HomeType,
		Condition:     prop.Condition,
		BrokerName:    prop.BrokerName,
		Description:   prop.Description,
		Active:        prop.IsActiveStatus,
		HasImages:     prop.HasImages,
		HasVideo:      prop.HasVideo,
		Has3DTour:     prop.Has3DTour,
		TotalParking:  prop.TotalParking,
		HasAC:         prop.HasAC,
		Appliance:     prop.Appliance,
		Heating:       prop.Heating,
		PetsAllowed:   prop.PetsAllowed,
	}, nil
}

func updatePropertyToDomain(ctx echo.Context, prop *propertyUpdateRequest) (*properties.Property, error) {
	idStr := ctx.Get("user_id")
	userID := idStr.(int)
	if userID == 0 {
		return nil, errors.New("invalid user id")
	}

	return &properties.Property{
		UserID:        userID,
		Location:      properties.Point{prop.Location.Lat, prop.Location.Long},
		Price:         prop.Price,
		PriceCurrency: prop.PriceCurrency,
		Address:       prop.FullAddress,
		Country:       prop.Address.Country,
		City:          prop.Address.City,
		State:         prop.Address.State,
		Street:        prop.Address.Street,
		ZipCode:       prop.Address.ZipCode,
		HouseNumber:   prop.Address.HouseNumber,
		Neighborhood:  prop.Address.Neighborhood,
		HasGarage:     prop.HasGarage,
		LivingSize:    prop.LivingSize,
		LotSize:       prop.LotSize,
		HomeSize:      prop.HomeSize,
		YearBuild:     prop.YearBuild,
		Bedroom:       prop.Bedroom,
		Bathroom:      prop.Bathroom,
		Floor:         prop.Floor,
		TotalFloors:   prop.TotalFloors,
		PropertyType:  prop.PropertyType,
		HomeType:      prop.HomeType,
		Condition:     prop.Condition,
		BrokerName:    prop.BrokerName,
		Description:   prop.Description,
		Active:        prop.IsActiveStatus,
		HasImages:     prop.HasImages,
		HasVideo:      prop.HasVideo,
		Has3DTour:     prop.Has3DTour,
		TotalParking:  prop.TotalParking,
		HasAC:         prop.HasAC,
		Appliance:     prop.Appliance,
		Heating:       prop.Heating,
		PetsAllowed:   prop.PetsAllowed,
	}, nil
}

func propertyToResponse(prop *properties.Property) *propertyResponse {
	propRe := &propertyResponse{
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
		LivingSize:     prop.LivingSize,
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
		Appliance:      prop.Appliance,
		Heating:        prop.Heating,
		PetsAllowed:    prop.PetsAllowed,

		ModelDateTime: common.ModelDateTime{
			CreatedAt: prop.CreatedAt.Unix(),
			UpdatedAt: prop.UpdatedAt.Unix(),
		},
	}

	for _, img := range prop.Images {
		propRe.Images = append(propRe.Images, img.ID.String())
	}

	return propRe
}

func listUserPropertiesToListRequest(request common.BaseSearchRequest) search.BaseFilters {
	return search.BaseFilters{
		Page: request.Page,
		Size: request.Size,
		Sort: request.Sort,
	}
}
