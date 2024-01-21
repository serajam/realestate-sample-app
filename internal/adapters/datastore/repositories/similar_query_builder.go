/* Copyright (C) Fedir Petryk */

package repositories

import (
	"fmt"

	"github.com/uptrace/bun"
)

type SimilarSearchQueryBuilder struct {
	query *bun.SelectQuery
}

func NewSimilarSearchQueryBuilder(query *bun.SelectQuery) *SimilarSearchQueryBuilder {
	return &SimilarSearchQueryBuilder{query: query}
}

func (q *SimilarSearchQueryBuilder) Query() *bun.SelectQuery {
	return q.query
}

func (q *SimilarSearchQueryBuilder) SetNotIds(ids []int) {
	q.query.Where("id NOT IN (?)", bun.In(ids))
}

func (q *SimilarSearchQueryBuilder) SetAreaWithin(p fmt.Stringer, distanceMeters int) {
	expr := fmt.Sprintf(`ST_DWithin(location::geography,ST_GeomFromText('%s',4326)::geography,%d)`, p.String(), distanceMeters)
	q.query = q.query.Where(expr)
}

func (q *SimilarSearchQueryBuilder) SetLimit(limit int) {
	q.query = q.query.Limit(limit)
}

func (q *SimilarSearchQueryBuilder) SetMinBedroom(count uint8) {
	q.query.Where("bedroom >= ?", count)
}

func (q *SimilarSearchQueryBuilder) SetPriceRange(min, max float32) {
	if min == 0 || max == 0 {
		return
	}

	q.query.Where("price BETWEEN ? AND ?", min, max)
}

func (q *SimilarSearchQueryBuilder) SetHomeSizeRange(min, max float32) {
	if min == 0 || max == 0 {
		return
	}

	q.query.Where("home_size BETWEEN ? AND ?", min, max)
}

func (q *SimilarSearchQueryBuilder) SetPropertyType(propertyType uint8) {
	if propertyType == 0 {
		return
	}

	q.query.Where("property_type = ?", propertyType)
}

// SetActive filters only active properties
func (q *SimilarSearchQueryBuilder) SetActive(active bool) {
	q.query.Where("active = ?", active)
}
