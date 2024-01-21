/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type SearchFilters struct {
	db *bun.DB
}

func NewSearchFilters(db *bun.DB) SearchFilters {
	return SearchFilters{db}
}

func (r SearchFilters) Create(ctx context.Context, search *domain.UserSearchFilters) error {
	_, err := r.db.NewInsert().Model(search).Returning("id").Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error creating search")
	}

	return err
}

func (r SearchFilters) Update(ctx context.Context, search *domain.UserSearchFilters) error {
	search.UpdatedAt = time.Now()
	_, err := r.db.NewUpdate().Model(search).WherePK().OmitZero().Where(
		"user_id = ?", search.UserID,
	).Returning("*").Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error updating search")
	}

	return nil
}

func (r SearchFilters) Delete(ctx context.Context, searchID, userID int) error {
	_, err := r.db.NewDelete().Model((*domain.UserSearchFilters)(nil)).Where(
		"id = ? AND user_id = ?", searchID, userID,
	).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "error deleting search")
	}

	return nil
}

func (r SearchFilters) List(ctx context.Context, userID int) ([]domain.UserSearchFilters, error) {
	var searches []domain.UserSearchFilters
	err := r.db.NewSelect().Model(&searches).Where("user_id = ?", userID).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error listing searches")
	}

	return searches, nil
}

func (r SearchFilters) Get(ctx context.Context, userID, searchID int) (*domain.UserSearchFilters, error) {
	var search domain.UserSearchFilters
	err := r.db.NewSelect().Model(&search).Where("id = ? AND user_id = ?", searchID, userID).Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error getting search")
	}

	return &search, nil
}
