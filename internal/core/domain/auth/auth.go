/* Copyright (C) Fedir Petryk */

package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/serajam/realestate-sample-app/internal/core/domain"
)

type Claims struct {
	Name      string `json:"name"`
	ID        int    `json:"id"`
	TokenUUID string `json:"token_uuid"`
	jwt.RegisteredClaims
}

type TokenGenerator struct {
	secret     string
	expiration int64
}

func NewTokenGenerator(secret string, expiration int64) TokenGenerator {
	return TokenGenerator{expiration: expiration, secret: secret}
}

func (a TokenGenerator) Generate(u *domain.User) (Token, error) {
	expiresIn := time.Now().Add(time.Duration(a.expiration) * time.Minute)
	tokenUUID := uuid.New().String()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, &Claims{
			Name:      u.Name,
			ID:        u.ID,
			TokenUUID: tokenUUID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiresIn),
			},
		},
	)

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return Token{}, err
	}

	return Token{
		Token:     tokenString,
		TokenUUID: tokenUUID,
		UserID:    u.ID,
		ExpiresIn: expiresIn.Unix(),
		Ttl:       time.Duration(a.expiration) * time.Minute,
	}, nil
}

func (a TokenGenerator) VerifyAndExtract(token string) (Claims, bool, error) {
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(
		token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(a.secret), nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Claims{}, false, errors.New("invalid token")
		}

		return Claims{}, false, err
	}

	if !tkn.Valid {
		return Claims{}, false, errors.New("invalid token")
	}

	return *claims, true, nil
}
