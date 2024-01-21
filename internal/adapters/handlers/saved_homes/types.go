/* Copyright (C) Fedir Petryk */

package saved_homes

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
)

type Service interface {
	Delete(ctx context.Context, home properties.UserSavedHome) error
	List(ctx context.Context, userId int, pagination domain.Pagination) ([]properties.Property, error)
	Add(ctx context.Context, home properties.UserSavedHome) (*properties.Property, error)
}
