/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/uptrace/bun"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type SavedHomes struct {
	db *bun.DB
}

func NewSavedHomes(db *bun.DB) SavedHomes {
	return SavedHomes{db}
}

func (r SavedHomes) Exists(ctx context.Context, home properties.UserSavedHome) (bool, error) {
	count, err := r.db.NewSelect().Model(&home).
		Where("user_id = ? AND property_id = ?", home.UserID, home.PropertyID).
		Count(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error checking if user saved home exists")
	}

	return count > 0, nil
}

func (r SavedHomes) Add(ctx context.Context, home properties.UserSavedHome) error {
	_, err := r.db.NewInsert().Model(&home).Exec(ctx)

	if err != nil {
		return errors.Wrap(err, "error saving property as saved home")
	}

	return nil
}

func (r SavedHomes) Delete(ctx context.Context, home properties.UserSavedHome) error {
	_, err := r.db.NewDelete().Model(&home).
		Where("user_id = ? AND property_id = ?", home.UserID, home.PropertyID).
		Exec(ctx)

	if err != nil {
		return errors.Wrap(err, "error deleting saved home")
	}

	return nil
}

func (r SavedHomes) List(
	ctx context.Context, userID int, pagination domain.Pagination,
) ([]properties.UserSavedHome, error) {
	var homes []properties.UserSavedHome
	err := r.db.NewSelect().Model(&homes).Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Size).
		Offset((pagination.Page - 1) * pagination.Size).
		Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error listing saved homes")
	}

	return homes, nil
}
