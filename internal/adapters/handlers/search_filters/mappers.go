/* Copyright (C) Fedir Petryk */

package search_filters

import "github.com/serajam/realestate-sample-app/internal/core/domain"

func mapSearchFiltersRequestToSearchFilters(req *searchFiltersRequest) (*domain.UserSearchFilters, error) {
	sf := &domain.UserSearchFilters{
		Name:           req.Name,
		EmailFrequency: req.EmailFrequency,
		Sort:           req.Sort,
		Subscribed:     req.Subscribed,
	}

	if req.Filters != nil {
		sf.Filters = &domain.SavedSearchesFilters{
			PriceFrom:      req.Filters.PriceFrom,
			PriceTo:        req.Filters.PriceTo,
			HomeSizeFrom:   req.Filters.HomeSizeFrom,
			HomeSizeTo:     req.Filters.HomeSizeTo,
			LotSizeFrom:    req.Filters.LotSizeFrom,
			LotSizeTo:      req.Filters.LotSizeTo,
			YearBuildFrom:  req.Filters.YearBuildFrom,
			YearBuildTo:    req.Filters.YearBuildTo,
			Bathroom:       req.Filters.Bathroom,
			BathroomExact:  req.Filters.BathroomExact,
			Bedroom:        req.Filters.Bedroom,
			BedroomExact:   req.Filters.BedroomExact,
			Condition:      req.Filters.Condition,
			HomeType:       req.Filters.HomeType,
			PropertyType:   req.Filters.PropertyType,
			HasAC:          req.Filters.HasAC,
			MustHaveGarage: req.Filters.MustHaveGarage,
			ParkingNumber:  req.Filters.ParkingNumber,
		}
	}

	if req.Polygon != nil {
		sf.Polygon = &domain.SearchPolygon{
			TopLat:     req.Polygon.TopLat,
			TopLong:    req.Polygon.TopLong,
			BottomLat:  req.Polygon.BottomLat,
			BottomLong: req.Polygon.BottomLong,
		}
	}

	return sf, nil
}

func mapSearchFiltersToSearchFiltersResponse(searchFilters []domain.UserSearchFilters) []searchFiltersResponse {
	var res []searchFiltersResponse

	for _, sf := range searchFilters {
		res = append(
			res, mapSearchFilterToSearchFilterResponse(&sf),
		)
	}

	return res
}

func mapSearchFilterToSearchFilterResponse(sf *domain.UserSearchFilters) searchFiltersResponse {
	return searchFiltersResponse{
		ID:             sf.ID,
		Name:           sf.Name,
		EmailFrequency: sf.EmailFrequency,
		Sort:           sf.Sort,
		Subscribed:     *sf.Subscribed,
		Filters: savedSearchFilters{
			PriceFrom:      sf.Filters.PriceFrom,
			PriceTo:        sf.Filters.PriceTo,
			HomeSizeFrom:   sf.Filters.HomeSizeFrom,
			HomeSizeTo:     sf.Filters.HomeSizeTo,
			LotSizeFrom:    sf.Filters.LotSizeFrom,
			LotSizeTo:      sf.Filters.LotSizeTo,
			YearBuildFrom:  sf.Filters.YearBuildFrom,
			YearBuildTo:    sf.Filters.YearBuildTo,
			Bathroom:       sf.Filters.Bathroom,
			BathroomExact:  sf.Filters.BathroomExact,
			Bedroom:        sf.Filters.Bedroom,
			BedroomExact:   sf.Filters.BedroomExact,
			Condition:      sf.Filters.Condition,
			HomeType:       sf.Filters.HomeType,
			PropertyType:   sf.Filters.PropertyType,
			HasAC:          sf.Filters.HasAC,
			MustHaveGarage: sf.Filters.MustHaveGarage,
			ParkingNumber:  sf.Filters.ParkingNumber,
		},
		Polygon: &polygon{
			TopLat:     sf.Polygon.TopLat,
			TopLong:    sf.Polygon.TopLong,
			BottomLat:  sf.Polygon.BottomLat,
			BottomLong: sf.Polygon.BottomLong,
		},
	}
}
