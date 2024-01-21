/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
)

type ProfileService interface {
	GetProfile(ctx context.Context, userID int) (*domain.User, error)
}

type PropertiesService interface {
	Get(ctx context.Context, id int) (*properties.Property, error)
	List(ctx context.Context, ids []int) ([]properties.Property, error)
}

type SimilarPropertiesService interface {
	Get(ctx context.Context, id int) (*properties.Property, error)
	List(ctx context.Context, ids []int) ([]properties.Property, error)
}

type ImagesService interface {
	DeleteAll(ctx context.Context, prop *properties.Property) error
	DeleteOne(ctx context.Context, propertyId int, imageId string) error
	Save(ctx context.Context, file *multipart.FileHeader, userId, propId int) error
	Get(ctx context.Context, id, imgType string) (*properties.Image, *minio.Object, error)
	ImageTypeExists(imgType string) bool
}
