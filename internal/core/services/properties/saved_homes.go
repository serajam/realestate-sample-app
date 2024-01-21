/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type SavedHomesService struct {
	savedHomesRepo       SavedHomesRepository
	propertiesRepo       PropertyRepository
	propertiesSearchRepo PropertySearchRepository
	logger               *zap.SugaredLogger
}

func NewSavedHomesSrv(
	savedHomesRepo SavedHomesRepository, propertiesRepo PropertyRepository,
	propertiesSearchRepo PropertySearchRepository, logger *zap.SugaredLogger,
) SavedHomesService {
	return SavedHomesService{
		savedHomesRepo:       savedHomesRepo,
		propertiesRepo:       propertiesRepo,
		propertiesSearchRepo: propertiesSearchRepo,
		logger:               logger,
	}
}

func (s SavedHomesService) List(ctx context.Context, userID int, pagination domain.Pagination) (
	[]properties.Property, error,
) {
	savedHomes, err := s.savedHomesRepo.List(ctx, userID, pagination)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.NotFound
		}

		s.logger.Errorw("SavedHomesService.List", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	var propertyIDs []int
	for _, home := range savedHomes {
		propertyIDs = append(propertyIDs, home.PropertyID)
	}

	if len(propertyIDs) == 0 {
		return nil, nil
	}

	queryBuilder := s.propertiesSearchRepo.SearchQueryBuilder()
	queryBuilder.SetIds(propertyIDs, true)

	properties, err := s.propertiesSearchRepo.Search(ctx, queryBuilder)
	if err != nil {
		s.logger.Errorw("SavedHomesService.List", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return properties, nil
}

func (s SavedHomesService) Delete(ctx context.Context, home properties.UserSavedHome) error {
	err := s.savedHomesRepo.Delete(ctx, home)
	if err != nil {
		s.logger.Errorw("SavedHomesService.DeleteAll", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
	}

	return nil
}

func (s SavedHomesService) Add(ctx context.Context, home properties.UserSavedHome) (*properties.Property, error) {
	exists, err := s.savedHomesRepo.Exists(ctx, home)
	if err != nil {
		s.logger.Errorw("SavedHomesService.Add", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	if exists {
		return nil, domainErrors.Exists
	}

	prop, err := s.propertiesRepo.Get(ctx, home.PropertyID)
	if err != nil {
		s.logger.Errorw("SavedHomesService.Add", "error", err)
		return nil, domainErrors.NotFound
	}

	err = s.savedHomesRepo.Add(ctx, home)
	if err != nil {
		s.logger.Errorw("SavedHomesService.Add", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailCreateOp}
	}

	return prop, nil
}
