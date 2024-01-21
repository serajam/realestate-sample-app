/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/uptrace/bun"
)

type PropertyImages struct {
	db *bun.DB
}

func NewPropertyImages(db *bun.DB) PropertyImages {
	return PropertyImages{db}
}

func (r PropertyImages) Create(ctx context.Context, image *properties.Image) error {
	_, err := r.db.NewInsert().Model(image).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r PropertyImages) Get(ctx context.Context, id string) (*properties.Image, error) {
	var image properties.Image
	err := r.db.NewSelect().Model(&image).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r PropertyImages) GetPropertyImage(ctx context.Context, propId int, id string) (*properties.Image, error) {
	var image properties.Image
	err := r.db.NewSelect().Model(&image).
		Where("id = ?", id).
		Where("property_id = ?", propId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r PropertyImages) GetUser(ctx context.Context, propertyId int, id string) (*properties.Image, error) {
	var image properties.Image
	err := r.db.NewSelect().Model(&image).
		Where("id = ?", id).
		Where("property_id = ?", propertyId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r PropertyImages) GetAll(ctx context.Context, id int) ([]properties.Image, error) {
	var images []properties.Image
	err := r.db.NewSelect().Model(&images).
		Where("property_id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// DeletePropertyImages deletes images from database
func (r PropertyImages) DeletePropertyImages(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().
		Model((*properties.Image)(nil)).
		Where("property_id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes image from database
func (r PropertyImages) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().
		Model((*properties.Image)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
