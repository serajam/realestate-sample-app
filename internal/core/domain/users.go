/* Copyright (C) Fedir Petryk */

package domain

import (
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type Email struct {
	Email string
}

type User struct {
	bun.BaseModel     `bun:"users"`
	ID                int       `bun:"id,pk,autoincrement"`
	AccountId         int       `bun:"account_id"`
	Name              string    `bun:"name"`
	Surname           string    `bun:"surname"`
	ScreenName        string    `bun:"screen_name"`
	CompanyName       string    `bun:"company_name"`
	PhotoUrl          string    `bun:"photo_url"`
	Phone             string    `bun:"phone"`
	AdditionalPhones  []string  `bun:"additional_phones,array"`
	Country           string    `bun:"country"`
	City              string    `bun:"city"`
	Zip               string    `bun:"zip"`
	Latitude          string    `bun:"latitude"`
	Longitude         string    `bun:"longitude"`
	Email             string    `bun:"email,unique"`
	Password          []byte    `bun:"password"`
	IsVerifiedEmail   bool      `bun:"is_verified_email"`
	IsCreatedPassword bool      `bun:"is_created_password"`
	IsVerifiedPhone   bool      `bun:"is_verified_phone"`
	IsGoogleLinked    bool      `bun:"is_google_linked"`
	IsFacebookLinked  bool      `bun:"is_facebook_linked"`
	IsAppleLinked     bool      `bun:"is_apple_linked"`
	Locked            bool      `bun:"locked"`
	Active            bool      `bun:"active"`
	CreatedAt         time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt         time.Time `bun:"updated_at,default:current_timestamp"`
}

// GenerateHashPassword hashes the given password using bcrypt
func (u *User) GenerateHashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

// CheckPasswordHash compares the given password with the hashed password
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	return err == nil
}
