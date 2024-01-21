/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"
	"fmt"

	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/uptrace/bun"
)

type SearchPropertyQueryBuilder interface {
	Query() *bun.SelectQuery
	SetArea(p fmt.Stringer)
	SetPaging(page, limit int)
	SetUserId(userID int)
	SetSort(sortType int)
	SetCity(city string)
	SetSortByIds(ids []int)
	SetIds(ids []int, required bool)

	SetBathroom(count uint8, exact bool)
	SetBedroom(count uint8, exact bool)

	SetPriceRange(min, max float32)
	SetLotSizeRange(min, max float32)
	SetHomeSizeRange(min, max float32)

	SetPropertyType(propertyType uint8)
	SetCondition(condition []int)
	SetHomeType(homeType []int)

	SetYearBuiltRange(min, max uint16)
	SetParking(parking *uint8)
	SetAC(hasAC *bool)

	SetActive(active bool)
}

type PropertySearchRepository struct {
	db *bun.DB
}

func NewPropertySearch(db *bun.DB) PropertySearchRepository {
	return PropertySearchRepository{db}
}

func (r PropertySearchRepository) SearchQueryBuilder() SearchPropertyQueryBuilder {
	return NewSearchQueryBuilder(r.db.NewSelect())
}

func (r PropertySearchRepository) Search(ctx context.Context, queryBuilder SearchPropertyQueryBuilder) (properties.Properties, error) {
	var properties properties.Properties
	q := queryBuilder.Query().Model(&properties).Relation("Images")

	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (r PropertySearchRepository) Count(ctx context.Context, q SearchPropertyQueryBuilder) (int, error) {
	count, err := q.Query().Offset(0).Limit(0).Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
