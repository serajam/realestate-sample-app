/* Copyright (C) Fedir Petryk */

package repositories

import (
	"fmt"

	"github.com/uptrace/bun"
)

type SearchQueryBuilder struct {
	query *bun.SelectQuery
}

func NewSearchQueryBuilder(query *bun.SelectQuery) *SearchQueryBuilder {
	return &SearchQueryBuilder{query: query}
}

func (q *SearchQueryBuilder) SetUserId(userID int) {
	q.query.Where("user_id = ?", userID)
}

func (q *SearchQueryBuilder) SetArea(p fmt.Stringer) {
	query := fmt.Sprintf(`ST_Contains(ST_GeomFromEWKT('%s'), p.location)`, p.String())
	q.query = q.query.Where(query)
}

func (q *SearchQueryBuilder) SetPaging(page, limit int) {
	q.query = q.query.Limit(limit)
	if page > 1 {
		q.query = q.query.Offset((page - 1) * limit)
	}
}

func (q *SearchQueryBuilder) SetIds(ids []int, required bool) {
	if len(ids) == 0 && !required {
		return
	}

	q.query.Where("id IN (?)", bun.In(ids))
}

func (q *SearchQueryBuilder) SetSortByIds(ids []int) {
	if len(ids) == 0 {
		return
	}

	q.query.OrderExpr("array_position(array[?], id)", bun.In(ids))
}

func (q *SearchQueryBuilder) SetSort(sort int) {
	switch sort {
	case newest:
		q.query.Order("created_at DESC")
	case priceAsc:
		q.query.Order("price ASC")
	case priceDesc:
		q.query.Order("price DESC")
	case bedroomAsc:
		q.query.Order("bedroom DESC")
	case bathroomAsc:
		q.query.Order("bathroom DESC")
	case homeSizeAsc:
		q.query.Order("home_size ASC")
	case lotSizeAsc:
		q.query.Order("lot_size ASC")
	}
}

func (q *SearchQueryBuilder) Query() *bun.SelectQuery {
	return q.query
}

func (q *SearchQueryBuilder) SetBathroom(count uint8, exact bool) {
	if count == 0 {
		return
	}

	if exact {
		q.query.Where("bathroom = ?", count)
	} else {
		q.query.Where("bathroom >= ?", count)
	}
}

func (q *SearchQueryBuilder) SetBedroom(count uint8, exact bool) {
	if count == 0 {
		return
	}

	if exact {
		q.query.Where("bedroom = ?", count)
	} else {
		q.query.Where("bedroom >= ?", count)
	}
}

func (q *SearchQueryBuilder) SetPriceRange(min, max float32) {
	if min == 0 || max == 0 {
		return
	}

	q.query.Where("price BETWEEN ? AND ?", min, max)
}

func (q *SearchQueryBuilder) SetLotSizeRange(min, max float32) {
	if min == 0 || max == 0 {
		return
	}

	q.query.Where("lot_size BETWEEN ? AND ?", min, max)
}

func (q *SearchQueryBuilder) SetHomeSizeRange(min, max float32) {
	if min == 0 || max == 0 {
		return
	}

	q.query.Where("home_size BETWEEN ? AND ?", min, max)
}

func (q *SearchQueryBuilder) SetPropertyType(propertyType uint8) {
	if propertyType == 0 {
		return
	}

	q.query.Where("property_type = ?", propertyType)
}

func (q *SearchQueryBuilder) SetCondition(condition []int) {
	if len(condition) == 0 {
		return
	}

	q.query.Where("condition IN (?)", bun.In(condition))
}

func (q *SearchQueryBuilder) SetHomeType(homeType []int) {
	if len(homeType) == 0 {
		return
	}

	q.query.Where("home_type IN (?)", bun.In(homeType))
}

func (q *SearchQueryBuilder) SetYearBuiltRange(min, max uint16) {
	if min == 0 || max == 0 {
		return
	}

	q.query.Where("year_build BETWEEN ? AND ?", min, max)
}

func (q *SearchQueryBuilder) SetParking(parking *uint8) {
	if parking == nil {
		return
	}

	q.query.Where("total_parking = ?", parking)
}

func (q *SearchQueryBuilder) SetAC(hasAC *bool) {
	if hasAC == nil {
		return
	}

	q.query.Where("has_ac = ?", hasAC)
}

// SetActive filters only active properties
func (q *SearchQueryBuilder) SetActive(active bool) {
	q.query.Where("active = ?", active)
}

func (q *SearchQueryBuilder) SetCity(city string) {
	if city == "" {
		return
	}

	q.query.Where("city = ?", city)
}
