/* Copyright (C) Fedir Petryk */

package properties

import (
	"context"

	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type SearchFilters struct {
	searchFiltersRepo SearchFiltersRepository
	logger            *zap.SugaredLogger
}

func NewSearchSrv(searchFiltersRepo SearchFiltersRepository, logger *zap.SugaredLogger) SearchFilters {
	return SearchFilters{searchFiltersRepo: searchFiltersRepo, logger: logger}
}

func (s SearchFilters) Create(ctx context.Context, searchFilters *domain.UserSearchFilters) (
	*domain.UserSearchFilters, error,
) {
	if err := s.searchFiltersRepo.Create(ctx, searchFilters); err != nil {
		s.logger.Errorw("error creating search filters", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailCreateOp}
	}

	return searchFilters, nil
}

func (s SearchFilters) Update(ctx context.Context, searchFilters *domain.UserSearchFilters) error {
	if err := s.searchFiltersRepo.Update(ctx, searchFilters); err != nil {
		s.logger.Errorw("error updating search filters", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailUpdateOp}
	}

	return nil
}

func (s SearchFilters) Delete(ctx context.Context, searchID, userID int) error {
	if err := s.searchFiltersRepo.Delete(ctx, searchID, userID); err != nil {
		s.logger.Errorw("error deleting search filters", "error", err)
		return domainErrors.OpFail{Op: domainErrors.MsgFailDeleteOp}
	}

	return nil
}

func (s SearchFilters) List(ctx context.Context, userID int) ([]domain.UserSearchFilters, error) {
	searches, err := s.searchFiltersRepo.List(ctx, userID)
	if err != nil {
		s.logger.Errorw("error listing search filters", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return searches, nil
}

func (s SearchFilters) Get(ctx context.Context, userID int, searchID int) (*domain.UserSearchFilters, error) {
	search, err := s.searchFiltersRepo.Get(ctx, userID, searchID)
	if err != nil {
		s.logger.Errorw("error getting search filters", "error", err)
		return nil, domainErrors.OpFail{Op: domainErrors.MsgFailGetOp}
	}

	return search, nil
}
