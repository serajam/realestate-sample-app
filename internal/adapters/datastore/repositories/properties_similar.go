/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"
	"fmt"

	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/uptrace/bun"
)

type SimilarPropertyQueryBuilder interface {
	Query() *bun.SelectQuery
	SetAreaWithin(p fmt.Stringer, distanceMeters int)

	SetNotIds(ids []int)
	SetLimit(num int)
	SetMinBedroom(count uint8)
	SetPriceRange(min, max float32)
	SetHomeSizeRange(min, max float32)
	SetPropertyType(propertyType uint8)
	SetActive(active bool)
}

type SimilarPropertySearchRepository struct {
	db *bun.DB
}

func NewSimilarPropertySearch(db *bun.DB) SimilarPropertySearchRepository {
	return SimilarPropertySearchRepository{db}
}

func (r SimilarPropertySearchRepository) SearchQueryBuilder() SimilarPropertyQueryBuilder {
	return NewSimilarSearchQueryBuilder(r.db.NewSelect())
}

func (r SimilarPropertySearchRepository) Search(ctx context.Context, queryBuilder SimilarPropertyQueryBuilder, originalLocation string) (
	properties.Properties,
	error,
) {
	var properties properties.Properties
	q := queryBuilder.Query().Model(&properties).Relation("Images")
	q.Where("street != ''")
	q.Where("house_number != ''")
	q.Column("*")
	q.ColumnExpr(
		` ST_Distance(
		        ST_GeomFromText(?, 4326),
		        location::geometry ) AS _distance`, originalLocation,
	).Order("_distance ASC")

	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (r SimilarPropertySearchRepository) Count(ctx context.Context, q SimilarPropertyQueryBuilder) (int, error) {
	count, err := q.Query().Offset(0).Limit(0).Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
