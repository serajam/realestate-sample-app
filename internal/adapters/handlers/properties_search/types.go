/* Copyright (C) Fedir Petryk */

package properties_search

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/serajam/realestate-sample-app/internal/core/domain/search"
)

type Service interface {
	Search(ctx context.Context, search *search.PropertySearchRequest) ([]properties.Property, int, error)
	SearchSimilar(ctx context.Context, id int) ([]properties.Property, error)
}
