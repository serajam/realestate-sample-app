/* Copyright (C) Fedir Petryk */

package user_properties

import (
	"context"
	"mime/multipart"

	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/serajam/realestate-sample-app/internal/core/domain/search"

	"github.com/minio/minio-go/v7"
	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type ProfileService interface {
	GetProfile(ctx context.Context, userID int) (*domain.User, error)
}

type UserPropertiesService interface {
	CreateUserProperty(ctx context.Context, prop *properties.Property) (*properties.Property, error)
	UpdateUserProperty(ctx context.Context, prop *properties.Property) (*properties.Property, error)
	DeleteUserProperty(ctx context.Context, prop *properties.Property) error
	ListByUser(ctx context.Context, userID int, filters search.BaseFilters) ([]properties.Property, error)
	GetUserProperty(ctx context.Context, id, userId int) (*properties.Property, error)
}

type ImagesService interface {
	DeleteAll(ctx context.Context, prop *properties.Property) error
	DeleteOne(ctx context.Context, propertyId int, imageId string) error
	Save(ctx context.Context, file *multipart.FileHeader, userId, propId int) error
	Get(ctx context.Context, id, imgType string) (*properties.Image, *minio.Object, error)
	ImageTypeExists(imgType string) bool
}
