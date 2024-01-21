/* Copyright (C) Fedir Petryk */

package users

import (
	"context"

	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type Profile struct {
	usersRepo UsersRepository
	propsRepo PropertyRepository

	publisher UsersPublisher

	logger *zap.SugaredLogger
}

func NewProfileSrv(
	repo UsersRepository,
	proprRepo PropertyRepository,
	publisher UsersPublisher,
	logger *zap.SugaredLogger,
) Profile {
	return Profile{usersRepo: repo, propsRepo: proprRepo, publisher: publisher, logger: logger}
}

func (s Profile) GetProfile(ctx context.Context, userID int) (*domain.User, error) {
	user, err := s.usersRepo.Get(ctx, userID)
	if err != nil {
		s.logger.Errorw("error getting user profile", "error", err)
		return nil, errors.OpFail{Op: errors.MsgFailGetOp}
	}

	return user, nil
}

func (s Profile) UpdateProfile(ctx context.Context, user *domain.User) error {
	err := s.usersRepo.Update(ctx, user)
	if err != nil {
		s.logger.Errorw("error update user profile", "error", err)
		return errors.OpFail{Op: errors.MsgFailUpdateOp}
	}

	return nil
}

func (s Profile) Deactivate(ctx context.Context, userID int) error {
	user, err := s.usersRepo.Get(ctx, userID)
	if err != nil {
		s.logger.Errorw("error getting user", "error", err)
		return errors.OpFail{Op: errors.MsgFailGetOp}
	}

	if !user.Active {
		return nil
	}

	err = s.usersRepo.UpdateActivity(ctx, userID, false)
	if err != nil {
		s.logger.Errorw("error deactivating user", "error", err)
		return errors.OpFail{Op: errors.MsgFailUpdateOp}
	}

	err = s.propsRepo.UpdateActivity(ctx, userID, false)
	if err != nil {
		s.logger.Errorw("error deactivating user properties", "error", err)
		return errors.OpFail{Op: errors.MsgFailUpdateOp}
	}

	go s.publisher.UserDeactivated(user.Email)

	return nil
}

func (s Profile) Activate(ctx context.Context, userID int) error {
	err := s.usersRepo.UpdateActivity(ctx, userID, true)
	if err != nil {
		s.logger.Errorw("error activating user", "error", err)
		return errors.OpFail{Op: errors.MsgFailUpdateOp}
	}

	err = s.propsRepo.UpdateActivity(ctx, userID, true)
	if err != nil {
		s.logger.Errorw("error activating user properties", "error", err)
		return errors.OpFail{Op: errors.MsgFailUpdateOp}
	}

	return nil
}
