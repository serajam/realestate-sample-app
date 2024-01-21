/* Copyright (C) Fedir Petryk */

package users

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type UsersRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Get(ctx context.Context, userID int) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	UpdateActivity(ctx context.Context, userID int, active bool) error
	Create(ctx context.Context, u domain.User, t domain.UserTokenAction) error
	EmailExists(ctx context.Context, email string) (bool, error)
	SetNewPwd(ctx context.Context, user *domain.User) error
}

type TokensRepository interface {
	Create(ctx context.Context, tokenAction *domain.UserTokenAction) error
	Update(ctx context.Context, tokenAction *domain.UserTokenAction) error
	GetByToken(ctx context.Context, token string) (*domain.UserTokenAction, error)
}

type Emailer interface {
	Send(subject, receiver string, body string) error
}

type UsersPublisher interface {
	UserDeactivated(email string)
	SignUp(receiver string, activationToken string)
}

type PropertyRepository interface {
	UpdateActivity(ctx context.Context, userID int, active bool) error
}
