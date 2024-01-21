/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"fmt"
	"net"

	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"go.uber.org/zap"

	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/search"
)

type PropertySearchService struct {
	propertiesRepo              PropertyRepository
	propertiesSearchRepo        PropertySearchRepository
	similarPropertiesSearchRepo SimilarPropertySearchRepository
	defaultPropSearchLimit      int

	logger *zap.SugaredLogger
}

func NewPropertySearchSrv(
	propertiesRepo PropertyRepository,
	propertiesSearchRepo PropertySearchRepository,
	similarPropertiesSearchRepo SimilarPropertySearchRepository,
	defaultPropLimit int,
	logger *zap.SugaredLogger,
) PropertySearchService {
	return PropertySearchService{
		propertiesRepo:              propertiesRepo,
		propertiesSearchRepo:        propertiesSearchRepo,
		similarPropertiesSearchRepo: similarPropertiesSearchRepo,
		defaultPropSearchLimit:      defaultPropLimit, logger: logger,
	}
}

func (s PropertySearchService) Search(ctx context.Context, search *search.PropertySearchRequest) (
	[]properties.Property, int, error,
) {
	var polygon fmt.Stringer
	if search.Polygon != nil {
		polygon = properties.NewPolygon(
			search.Polygon.TopLat, search.Polygon.TopLong, search.Polygon.BottomLat, search.Polygon.BottomLong,
		)
	}

	queryBuilder := s.propertiesSearchRepo.SearchQueryBuilder()
	if polygon != nil {
		queryBuilder.SetArea(polygon)
	}

	if search.City != "" {
		queryBuilder.SetCity(search.City)
	}

	queryBuilder.SetActive(true)
	queryBuilder.SetSort(search.Sort)
	queryBuilder.SetPriceRange(search.PriceFrom, search.PriceTo)
	queryBuilder.SetLotSizeRange(search.LotSizeFrom, search.LotSizeTo)
	queryBuilder.SetHomeSizeRange(search.HomeSizeFrom, search.HomeSizeTo)
	queryBuilder.SetAC(search.HasAC)
	queryBuilder.SetParking(search.ParkingNumber)
	queryBuilder.SetYearBuiltRange(search.YearBuildFrom, search.YearBuildTo)
	queryBuilder.SetPropertyType(search.PropertyType)
	queryBuilder.SetCondition(search.Condition)
	queryBuilder.SetHomeType(search.HomeType)
	queryBuilder.SetBathroom(search.Bathroom, search.BathroomExact)
	queryBuilder.SetBedroom(search.Bedroom, search.BedroomExact)

	limit := s.defaultPropSearchLimit
	if search.Size > 0 {
		limit = search.Size
	}
	queryBuilder.SetPaging(search.Page, limit)

	properties, err := s.propertiesSearchRepo.Search(ctx, queryBuilder)
	if err != nil {
		s.logger.Errorw("error searching properties", "error", err)
		return nil, 0, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	count, err := s.propertiesSearchRepo.Count(ctx, queryBuilder)
	if err != nil {
		s.logger.Errorw("error counting properties", "error", err)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, 0, domainErrors.RequestTimeout
		}

		return nil, 0, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return properties, count, nil
}

func (s PropertySearchService) SearchSimilar(ctx context.Context, id int) (
	[]properties.Property, error,
) {
	originProperty, err := s.propertiesRepo.Get(ctx, id)
	if err != nil {
		s.logger.Errorw("error getting origin property for similar search", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	distanceMeters := 4000

	queryBuilder := s.similarPropertiesSearchRepo.SearchQueryBuilder()
	queryBuilder.SetNotIds([]int{id})
	queryBuilder.SetAreaWithin(originProperty.Location, distanceMeters)
	queryBuilder.SetLimit(6)
	queryBuilder.SetMinBedroom(originProperty.Bedroom)
	queryBuilder.SetPriceRange(originProperty.Price*0.8, originProperty.Price*1.2)
	queryBuilder.SetHomeSizeRange(originProperty.HomeSize*0.8, originProperty.HomeSize*1.2)
	queryBuilder.SetPropertyType(originProperty.PropertyType)
	queryBuilder.SetActive(true)

	properties, err := s.similarPropertiesSearchRepo.Search(ctx, queryBuilder, originProperty.Location.String())
	if err != nil {
		s.logger.Errorw("error searching similar properties", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	for i := 0; i < len(properties); i++ {
		if properties[i].Images == nil {
			continue
		}
		properties[i].Images = properties[i].Images[:1]
	}

	return properties, nil
}
