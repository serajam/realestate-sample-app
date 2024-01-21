/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"github.com/serajam/realestate-sample-app/internal/core/domain/search"
	"go.uber.org/zap"

	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type UserPropertyService struct {
	propertiesRepo       PropertyRepository
	propertiesSearchRepo PropertySearchRepository

	defaultPropSearchLimit int
	logger                 *zap.SugaredLogger
}

func NewUserPropertySrv(
	propertiesRepo PropertyRepository,
	propertiesSearchRepo PropertySearchRepository,
	defaultPropSearchLimit int,
	logger *zap.SugaredLogger,
) UserPropertyService {
	return UserPropertyService{
		propertiesRepo:         propertiesRepo,
		propertiesSearchRepo:   propertiesSearchRepo,
		logger:                 logger,
		defaultPropSearchLimit: defaultPropSearchLimit,
	}
}

func (s UserPropertyService) ListByUser(ctx context.Context, userID int, filters search.BaseFilters) ([]properties.Property, error) {
	limit := s.defaultPropSearchLimit
	if filters.Size > 0 {
		limit = filters.Size
	}

	queryBuilder := s.propertiesSearchRepo.SearchQueryBuilder()
	queryBuilder.SetPaging(filters.Page, limit)
	queryBuilder.SetUserId(userID)

	properties, err := s.propertiesSearchRepo.Search(ctx, queryBuilder)
	if err != nil {
		s.logger.Errorw("error searching properties", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return properties, nil
}

func (s UserPropertyService) GetUserProperty(ctx context.Context, id, userId int) (*properties.Property, error) {
	properties, err := s.propertiesRepo.GetByUser(ctx, id, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.NotFound
		}

		s.logger.Errorw("error getting property", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return properties, nil
}

func (s UserPropertyService) CreateUserProperty(ctx context.Context, prop *properties.Property) (*properties.Property, error) {
	err := s.propertiesRepo.Create(ctx, prop)
	if err != nil {
		s.logger.Errorw("error creating property", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailCreateOp}
	}

	return prop, nil
}

func (s UserPropertyService) UpdateUserProperty(ctx context.Context, prop *properties.Property) (*properties.Property, error) {
	_, err := s.propertiesRepo.GetByUser(ctx, prop.ID, prop.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.NotFound
		}
		s.logger.Errorw("error getting property", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	prop.UpdatedAt = time.Now()

	err = s.propertiesRepo.Update(ctx, prop)
	if err != nil {
		s.logger.Errorw("error updating property", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailUpdateOp}
	}

	return prop, nil
}

func (s UserPropertyService) DeleteUserProperty(ctx context.Context, prop *properties.Property) error {
	_, err := s.propertiesRepo.GetByUser(ctx, prop.ID, prop.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.NotFound
		}
		s.logger.Errorw("error getting property", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	err = s.propertiesRepo.Delete(ctx, prop)
	if err != nil {
		s.logger.Errorw("error deleting property", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
	}

	return nil
}
