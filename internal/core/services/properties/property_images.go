/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"database/sql"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"go.uber.org/zap"
)

const (
	ImageTypeOrigin    = "origin"
	ImageTypeThumbnail = "thumb"
)

type ImageService struct {
	imgStorage   ImageStorage
	imagesRepo   ImagesRepository
	imgProcessor ImageProcessing
	imageTypes   map[string]struct{}
	logger       *zap.SugaredLogger
}

func NewImageService(
	imgStorage ImageStorage,
	imagesRepo ImagesRepository,
	imgProcessor ImageProcessing,
	logger *zap.SugaredLogger,
) *ImageService {
	return &ImageService{
		imgStorage:   imgStorage,
		imagesRepo:   imagesRepo,
		imgProcessor: imgProcessor,
		imageTypes:   map[string]struct{}{ImageTypeOrigin: {}, ImageTypeThumbnail: {}},
		logger:       logger,
	}
}

func (s ImageService) ImageTypeExists(imgType string) bool {
	_, ok := s.imageTypes[imgType]
	return ok
}

func (s ImageService) Save(ctx context.Context, file *multipart.FileHeader, userId, propId int) error {
	image, err := s.imgProcessor.Process(file)
	if err != nil {
		s.logger.Errorw("error processing property image", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailCreateOp}
	}

	image.UserID = userId
	image.PropertyID = propId

	err = s.imgStorage.Upload(ctx, image)
	if err != nil {
		s.logger.Errorw("error uploading property image to storage", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailCreateOp}
	}

	err = s.imagesRepo.Create(ctx, image)
	if err != nil {

		for t := range s.imageTypes {
			err = s.imgStorage.Del(ctx, properties.ImageObjectName(image.ID.String(), t))
			if err != nil {
				s.logger.Errorw("error deleting property image from storage", "error", err)
			}
		}

		s.logger.Errorw("error creating property image in db", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailCreateOp}
	}

	return nil
}

func (s ImageService) Get(ctx context.Context, id, imgType string) (*properties.Image, *minio.Object, error) {
	img, err := s.imagesRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, domainErrors.NotFound
		}
		s.logger.Errorw("error getting property image", "error", err)
		return nil, nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	obj, err := s.imgStorage.Get(ctx, properties.ImageObjectName(img.ID.String(), imgType))
	if err != nil {
		return nil, nil, err
	}

	return img, obj, nil
}

func (s ImageService) DeleteAll(ctx context.Context, prop *properties.Property) error {
	imgs, err := s.imagesRepo.GetAll(ctx, prop.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.NotFound
		}
		s.logger.Errorw("error getting property images", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	for t := range s.imageTypes {
		for _, img := range imgs {
			err = s.imgStorage.Del(ctx, properties.ImageObjectName(img.ID.String(), t))
			if err != nil {
				s.logger.Errorw("error deleting property image from storage", "error", err)
				return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
			}
		}
	}

	err = s.imagesRepo.DeletePropertyImages(ctx, prop.ID)
	if err != nil {
		s.logger.Errorw("error deleting property images from database", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
	}

	return nil
}

func (s ImageService) DeleteOne(ctx context.Context, propertyId int, imageId string) error {
	img, err := s.imagesRepo.GetPropertyImage(ctx, propertyId, imageId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.NotFound
		}
		s.logger.Errorw("error getting property images", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	for t := range s.imageTypes {
		err = s.imgStorage.Del(ctx, properties.ImageObjectName(img.ID.String(), t))
		if err != nil {
			s.logger.Errorw("error deleting property image from storage", "error", err)
			return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
		}
	}

	err = s.imagesRepo.Delete(ctx, imageId)
	if err != nil {
		s.logger.Errorw("error deleting property images from database", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
	}

	return nil
}
