/* Copyright (C) Fedir Petryk */

package search_filters

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type Service interface {
	Create(ctx context.Context, req *domain.UserSearchFilters) (*domain.UserSearchFilters, error)
	Update(ctx context.Context, req *domain.UserSearchFilters) error
	Delete(ctx context.Context, searchID, userID int) error
	List(ctx context.Context, userID int) ([]domain.UserSearchFilters, error)
	Get(ctx context.Context, userID int, searchID int) (*domain.UserSearchFilters, error)
}
