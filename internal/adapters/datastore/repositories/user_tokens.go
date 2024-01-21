/* Copyright (C) Fedir Petryk */

package repositories

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type UserTokenActionRepo struct {
	db *bun.DB
}

func NewUserTokenAction(db *bun.DB) UserTokenActionRepo {
	return UserTokenActionRepo{db: db}
}

func (r UserTokenActionRepo) Create(ctx context.Context, tokenAction *domain.UserTokenAction) error {
	_, err := r.db.NewInsert().Model(tokenAction).Exec(ctx)
	return err
}

func (r UserTokenActionRepo) Delete(ctx context.Context, token string) error {
	_, err := r.db.NewDelete().
		Model((*domain.UserTokenAction)(nil)).
		Where("token = ?", token).
		Exec(ctx)
	return err
}

func (r UserTokenActionRepo) Update(ctx context.Context, tokenAction *domain.UserTokenAction) error {
	_, err := r.db.NewUpdate().
		Model(tokenAction).
		WherePK().
		Exec(ctx)
	return err
}

func (r UserTokenActionRepo) GetByToken(ctx context.Context, token string) (*domain.UserTokenAction, error) {
	var tokenAction domain.UserTokenAction
	err := r.db.NewSelect().
		Model(&tokenAction).
		Where("token = ?", token).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &tokenAction, nil
}
