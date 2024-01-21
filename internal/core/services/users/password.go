/* Copyright (C) Fedir Petryk */

package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	domainErrors "github.com/serajam/realestate-sample-app/internal/core/domain/errors"
)

type Password struct {
	userRepository          UsersRepository
	tokenRepository         TokensRepository
	emailer                 Emailer
	logger                  *zap.SugaredLogger
	pwdResetExpirationHours int
	pwdResetUrl             string
}

func NewPasswordSrv(
	signRepository UsersRepository,
	tokenRepository TokensRepository,
	emailer Emailer,
	pwdResetExpirationHours int,
	pwdResetUrl string,
	logger *zap.SugaredLogger,
) Password {
	return Password{
		userRepository:          signRepository,
		emailer:                 emailer,
		logger:                  logger,
		tokenRepository:         tokenRepository,
		pwdResetUrl:             pwdResetUrl,
		pwdResetExpirationHours: pwdResetExpirationHours,
	}
}

func (s Password) SetNewPassword(ctx context.Context, token, newPwd string) error {
	tokenModel, err := s.tokenRepository.GetByToken(ctx, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.Dummy
		}

		s.logger.Errorw(err.Error(), "method", "SetNewPassword")
		return domainErrors.Internal
	}

	if tokenModel == nil {
		return domainErrors.Dummy
	}

	payload := domain.Email{}
	err = json.Unmarshal(tokenModel.Payload, &payload)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SetNewPassword")
		return domainErrors.Internal
	}

	user, err := s.userRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SetNewPassword")
		return domainErrors.Internal
	}

	user.Password = []byte(newPwd)

	err = user.GenerateHashPassword()
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SetNewPassword")
		return domainErrors.Internal
	}

	err = s.userRepository.SetNewPwd(ctx, user)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SetNewPassword")
		return domainErrors.Internal
	}

	err = s.emailer.Send(
		"Password Reset", payload.Email, fmt.Sprintf("Passowrd was reset successfully"),
	)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "SetNewPassword")
	}

	return nil
}

func (s Password) ResetPassword(ctx context.Context, email string) error {
	exists, err := s.userRepository.EmailExists(ctx, email)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "ResetPassword")
		return domainErrors.Internal
	}

	if !exists {
		return domainErrors.Dummy
	}

	emailModel := domain.Email{Email: email}
	payload, err := json.Marshal(emailModel)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "ResetPassword")
		return domainErrors.Internal
	}

	token := domain.UserTokenAction{
		Token:       domain.RandomToken(),
		Action:      domain.PasswordReset,
		Payload:     payload,
		TokenExpiry: time.Now().Add(time.Hour * time.Duration(s.pwdResetExpirationHours)),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err = s.tokenRepository.Create(ctx, &token)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "ResetPassword")
		return domainErrors.Internal
	}

	err = s.emailer.Send(
		"Password Reset", email, fmt.Sprintf(`Password reset link: %s`, s.pwdResetUrl+"/"+token.Token),
	)
	if err != nil {
		s.logger.Errorw(err.Error(), "method", "ResetPassword")
	}

	return nil
}
