/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/uptrace/bun"
)

type PropertyRepository struct {
	db *bun.DB
}

func NewProperty(db *bun.DB) PropertyRepository {
	return PropertyRepository{db}
}

func (r PropertyRepository) Create(ctx context.Context, property *properties.Property) error {
	_, err := r.db.NewInsert().Model(property).Returning("id, created_at, updated_at").Exec(ctx)

	if err != nil {
		return errors.Wrap(err, "error creating property")
	}

	return nil
}

func (r PropertyRepository) Update(ctx context.Context, property *properties.Property) error {
	property.UpdatedAt = time.Now()
	_, err := r.db.NewUpdate().Model(property).OmitZero().WherePK().Returning("*").Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error updating property")
	}

	return nil
}

func (r PropertyRepository) Delete(ctx context.Context, property *properties.Property) error {
	_, err := r.db.NewDelete().
		Model((*properties.Property)(nil)).
		Where("id = ?", property.ID).
		Where("user_id = ?", property.UserID).
		Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error deleting property")
	}

	return nil
}

func (r PropertyRepository) DeleteAll(ctx context.Context) error {
	_, err := r.db.NewRaw("TRUNCATE TABLE properties CASCADE").Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error deleting property")
	}

	return nil
}

func (r PropertyRepository) Get(ctx context.Context, id int) (*properties.Property, error) {
	property := new(properties.Property)
	err := r.db.NewSelect().
		Model(property).
		Relation("Images").
		Where("p.id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting property")
	}

	return property, nil
}

func (r PropertyRepository) GetByUser(ctx context.Context, id, userID int) (*properties.Property, error) {
	property := new(properties.Property)
	err := r.db.NewSelect().
		Model(property).
		Relation("Images").
		Where("id = ? AND user_id = ?", id, userID).
		Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting property")
	}

	return property, nil
}

func (r PropertyRepository) UpdateActivity(ctx context.Context, userID int, active bool) error {
	_, err := r.db.NewUpdate().
		Model(&properties.Property{}).
		Set("active = ?", active).
		Where("user_id = ?", userID).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error deactivating user properties")
	}

	return nil
}

func (r PropertyRepository) SearchQueryBuilder() SearchPropertyQueryBuilder {
	return NewSearchQueryBuilder(r.db.NewSelect().Model(&properties.Property{}))
}
