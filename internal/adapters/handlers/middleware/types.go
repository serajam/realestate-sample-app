/* Copyright (C) Fedir Petryk */

package middleware

import (
	"context"

	"github.com/serajam/realestate-sample-app/internal/core/domain/auth"
)

type AuthService interface {
	Validate(ctx context.Context, token string) (auth.Token, error)
}
