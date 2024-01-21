/* Copyright (C) Fedir Petryk */

package users

import "github.com/serajam/realestate-sample-app/internal/core/domain"

func mapSignUpRequestToUserDomain(req signUpRequest) domain.User {
	return domain.User{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: []byte(req.Password),
	}
}

func mapProfileRequestToUserDomain(req *profileRequest) *domain.User {
	return &domain.User{
		Name:             req.Name,
		Surname:          req.Surname,
		ScreenName:       req.ScreenName,
		PhotoUrl:         req.PhotoUrl,
		CompanyName:      req.CompanyName,
		Phone:            req.Phone,
		AdditionalPhones: req.AdditionalPhones,
		Country:          req.Country,
		City:             req.City,
		Zip:              req.Zip,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
	}
}

func mapUserDomainToProfile(user *domain.User) *profileResponse {
	return &profileResponse{
		AccountId:         user.AccountId,
		Name:              user.Name,
		Surname:           user.Surname,
		ScreenName:        user.ScreenName,
		CompanyName:       user.CompanyName,
		PhotoUrl:          user.PhotoUrl,
		Phone:             user.Phone,
		AdditionalPhones:  user.AdditionalPhones,
		Country:           user.Country,
		City:              user.City,
		Zip:               user.Zip,
		Latitude:          user.Latitude,
		Longitude:         user.Longitude,
		Email:             user.Email,
		IsVerifiedEmail:   user.IsVerifiedEmail,
		IsCreatedPassword: user.IsCreatedPassword,
		IsVerifiedPhone:   user.IsVerifiedPhone,
		IsGoogleLinked:    user.IsGoogleLinked,
		IsFacebookLinked:  user.IsFacebookLinked,
		IsAppleLinked:     user.IsAppleLinked,
		IsAccountLocked:   user.Locked,
	}
}
