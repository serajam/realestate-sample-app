/* Copyright (C) Fedir Petryk */

package users

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
	"github.com/serajam/realestate-sample-app/internal/core/domain/auth"
)

type ProfileService interface {
	GetProfile(ctx context.Context, userID int) (*domain.User, error)
	UpdateProfile(ctx context.Context, user *domain.User) error
	Deactivate(ctx context.Context, userID int) error
}

type SignUpService interface {
	SignUp(ctx context.Context, user domain.User) error
}

type AuthService interface {
	Authenticate(ctx context.Context, email, pwd, deviceID string) (
		auth.AccessToken, auth.RefreshToken, error,
	)
	SignOut(ctx context.Context, userID int, deviceID string) error
	SignOutAll(ctx context.Context, userID int) error
	Refresh(ctx context.Context, userID int, tokenUUID, deviceID string) (
		auth.AccessToken, auth.RefreshToken, error,
	)
}

type ResetPwdService interface {
	ResetPassword(ctx context.Context, email string) error
	SetNewPassword(ctx context.Context, token, newPwd string) error
}
