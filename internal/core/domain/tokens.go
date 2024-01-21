/* Copyright (C) Fedir Petryk */

package domain

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
)

type TokenActionType string

const (
	EmailVerification TokenActionType = "email_verification"
	PasswordReset     TokenActionType = "password_reset"
)

type UserTokenAction struct {
	bun.BaseModel `bun:"user_token_actions"`
	ID            int             `bun:"id,pk,autoincrement"`
	Token         string          `bun:"token"`
	UserID        *int            `bun:"user_id,nullzero"`
	Action        TokenActionType `bun:"action"`
	Payload       json.RawMessage `bun:"payload,nullzero"`
	TokenExpiry   time.Time       `bun:"token_expiry"`
	CreatedAt     time.Time       `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time       `bun:"updated_at,default:current_timestamp"`
}

func (t UserTokenAction) IsExpired() bool {
	return time.Now().After(t.TokenExpiry)
}

func RandomToken() string {
	b := make([]byte, 34)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Read(b)

	return fmt.Sprintf("%x", b)[2:34]
}
