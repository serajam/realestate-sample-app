/* Copyright (C) Fedir Petryk */

package users

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type SignUp struct {
	signupRepository            UsersRepository
	logger                      *zap.SugaredLogger
	publisher                   UsersPublisher
	registrationExpirationHours int
}

func NewSignupSrv(
	signRepository UsersRepository,
	publisher UsersPublisher,
	registrationExpirationHours int,
	logger *zap.SugaredLogger,
) SignUp {
	return SignUp{
		signupRepository:            signRepository,
		logger:                      logger,
		publisher:                   publisher,
		registrationExpirationHours: registrationExpirationHours,
	}
}

func (s SignUp) SignUp(ctx context.Context, user domain.User) error {
	exists, err := s.signupRepository.EmailExists(ctx, user.Email)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SignUp")
		return errors.User{Msg: errors.MsgUserSingupFailed}
	}

	if exists {
		return errors.User{Msg: errors.MsgEmailExists}
	}

	err = user.GenerateHashPassword()
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SignUp")
		return errors.User{Msg: errors.MsgUserSingupFailed}
	}

	token := domain.UserTokenAction{
		Token:       domain.RandomToken(),
		Action:      domain.EmailVerification,
		TokenExpiry: time.Now().Add(time.Hour * time.Duration(s.registrationExpirationHours)),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err = s.signupRepository.Create(ctx, user, token)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SignUp")
		return errors.User{Msg: errors.MsgUserSingupFailed}
	}

	go s.publisher.SignUp(user.Email, token.Token)

	return nil
}
