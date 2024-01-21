/* Copyright (C) Fedir Petryk */

package users

type signUpRequest struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=30"`
	Surname  string `json:"surname" validate:"omitempty,min=2,max=30"`
	Email    string `json:"email" validate:"required,email,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
} // @name signUpRequest

type signIn struct {
	Email    string `json:"email" validate:"required,email,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
} // @name signIn

type authToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
} // @name authToken

// resetPwdRequest is the request DTO for resetting a user's password.
type resetPwdRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
} // @name resetPwdRequest

// newPasswordRequest is the request DTO for setting a new password.
type newPasswordRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Token    string `param:"token" validate:"required,len=32"`
} // @name newPasswordRequest

type profileResponse struct {
	AccountId            int      `json:"accountId,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Surname              string   `json:"surname,omitempty"`
	ScreenName           string   `json:"screenName,omitempty"`
	CompanyName          string   `json:"companyName,omitempty"`
	PhotoUrl             string   `json:"photoUrl,omitempty"`
	Phone                string   `json:"phone,omitempty"`
	AdditionalPhones     []string `json:"additionalPhones,omitempty"`
	Country              string   `json:"country,omitempty"`
	City                 string   `json:"city,omitempty"`
	Zip                  string   `json:"zip,omitempty"`
	Latitude             string   `json:"latitude,omitempty"`
	Longitude            string   `json:"longitude,omitempty"`
	Email                string   `json:"email,unique,omitempty"`
	IsVerifiedEmail      bool     `json:"isVerifiedEmail,omitempty"`
	IsCreatedPassword    bool     `json:"isCreatedPassword,omitempty"`
	IsVerifiedPhone      bool     `json:"isVerifiedPhone,omitempty"`
	IsGoogleLinked       bool     `json:"isGoogleLinked,omitempty"`
	IsFacebookLinked     bool     `json:"isFacebookLinked,omitempty"`
	IsAppleLinked        bool     `json:"isAppleLinked,omitempty"`
	IsAccountLocked      bool     `json:"isAccountLocked,omitempty"`
	IsAccountDeactivated bool     `json:"isAccountDeactivated,omitempty"`
} // @name profileResponse

type profileRequest struct {
	Name             string   `json:"name,omitempty" validate:"omitempty,min=2,max=30"`
	Surname          string   `json:"surname,omitempty" validate:"omitempty,min=2,max=30"`
	ScreenName       string   `json:"screenName,omitempty" validate:"omitempty,min=2,max=30"`
	CompanyName      string   `json:"companyName,omitempty" validate:"omitempty,min=2,max=200"`
	PhotoUrl         string   `json:"photoUrl,omitempty" validate:"omitempty,min=2,max=200"`
	Phone            string   `json:"phone,omitempty" validate:"omitempty,min=2,max=30"`
	AdditionalPhones []string `json:"additionalPhones,omitempty" validate:"omitempty,min=2,max=30"`
	Country          string   `json:"country,omitempty" validate:"omitempty,min=2,max=30"`
	City             string   `json:"city,omitempty" validate:"omitempty,min=2,max=30"`
	Zip              string   `json:"zip,omitempty" validate:"omitempty,min=2,max=15"`
	Latitude         string   `json:"latitude,omitempty" validate:"omitempty,min=2,max=15"`
	Longitude        string   `json:"longitude,omitempty" validate:"omitempty,min=2,max=15"`
} // @name profileRequest
