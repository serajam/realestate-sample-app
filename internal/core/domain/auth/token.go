/* Copyright (C) Fedir Petryk */

package auth

import (
	"time"
)

type (
	AccessToken  string
	RefreshToken string
	DeviceId     string
)

type Token struct {
	Token     string
	TokenUUID string
	UserID    int
	ExpiresIn int64
	Ttl       time.Duration
}

type TokenMeta struct {
	UserID    int
	TokenType string
	DeviceId  string
	Token     string
}

func FromClaims(claims Claims) Token {
	return Token{
		UserID:    claims.ID,
		ExpiresIn: claims.ExpiresAt.Time.Unix(),
		TokenUUID: claims.TokenUUID,
	}
}
