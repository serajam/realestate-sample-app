/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
	"go.uber.org/zap"

	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type PropertyService struct {
	propertiesRepo       PropertyRepository
	propertiesSearchRepo PropertySearchRepository

	logger *zap.SugaredLogger
}

func NewPropertySrv(
	propertiesRepo PropertyRepository,
	propertiesSearchRepo PropertySearchRepository,
	logger *zap.SugaredLogger,
) PropertyService {
	return PropertyService{
		propertiesRepo:       propertiesRepo,
		propertiesSearchRepo: propertiesSearchRepo,
		logger:               logger,
	}
}

func (s PropertyService) List(ctx context.Context, ids []int) ([]properties.Property, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	queryBuilder := s.propertiesSearchRepo.SearchQueryBuilder()
	queryBuilder.SetIds(ids, false)
	queryBuilder.SetActive(true)
	queryBuilder.SetSortByIds(ids)

	properties, err := s.propertiesSearchRepo.Search(ctx, queryBuilder)
	if err != nil {
		s.logger.Errorw("error searching properties", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return properties, nil
}

func (s PropertyService) Get(ctx context.Context, id int) (*properties.Property, error) {
	properties, err := s.propertiesRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.NotFound
		}

		s.logger.Errorw("error getting property", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	if *properties.Active == false {
		return nil, domainErrors.NotFound
	}

	return properties, nil
}
