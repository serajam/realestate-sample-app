/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/serajam/realestate-sample-app/internal/adapters/datastore/repositories"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type SavedHomesRepository interface {
	Exists(ctx context.Context, home properties.UserSavedHome) (bool, error)
	Add(ctx context.Context, home properties.UserSavedHome) error
	Delete(ctx context.Context, home properties.UserSavedHome) error
	List(ctx context.Context, userID int, pagination domain.Pagination) ([]properties.UserSavedHome, error)
}

type Searchable interface {
	SearchQueryBuilder() repositories.SearchPropertyQueryBuilder
}

type SimilarSearchable interface {
	SearchQueryBuilder() repositories.SimilarPropertyQueryBuilder
}

type PropertyRepository interface {
	Get(ctx context.Context, id int) (*properties.Property, error)
	GetByUser(ctx context.Context, id, userID int) (*properties.Property, error)
	Create(ctx context.Context, property *properties.Property) error
	Update(ctx context.Context, property *properties.Property) error
	Delete(ctx context.Context, property *properties.Property) error
	UpdateActivity(ctx context.Context, userID int, active bool) error

	Searchable
}

type SimilarPropertySearchRepository interface {
	Search(ctx context.Context, q repositories.SimilarPropertyQueryBuilder, location string) (properties.Properties, error)
	Count(ctx context.Context, q repositories.SimilarPropertyQueryBuilder) (int, error)

	SimilarSearchable
}

type PropertySearchRepository interface {
	Search(ctx context.Context, q repositories.SearchPropertyQueryBuilder) (properties.Properties, error)
	Count(ctx context.Context, q repositories.SearchPropertyQueryBuilder) (int, error)

	Searchable
}

type SearchFiltersRepository interface {
	Create(ctx context.Context, search *domain.UserSearchFilters) error
	Update(ctx context.Context, search *domain.UserSearchFilters) error
	Delete(ctx context.Context, searchID, userID int) error
	List(ctx context.Context, userID int) ([]domain.UserSearchFilters, error)
	Get(ctx context.Context, userID, searchID int) (*domain.UserSearchFilters, error)
}

type ImageStorage interface {
	Upload(ctx context.Context, img *properties.Image) error
	Get(ctx context.Context, name string) (*minio.Object, error)
	Del(ctx context.Context, name string) error
}

type ImagesRepository interface {
	Create(ctx context.Context, image *properties.Image) error
	Get(ctx context.Context, id string) (*properties.Image, error)
	GetPropertyImage(ctx context.Context, propId int, id string) (*properties.Image, error)
	GetAll(ctx context.Context, id int) ([]properties.Image, error)
	Delete(ctx context.Context, id string) error
	DeletePropertyImages(ctx context.Context, id int) error
}

type ImageProcessing interface {
	Process(file *multipart.FileHeader) (*properties.Image, error)
}
