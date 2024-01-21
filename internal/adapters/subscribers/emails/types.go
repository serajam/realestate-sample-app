/* Copyright (C) Fedir Petryk */

package user

type EmailService interface {
	UserDeactivated(receiver string) error
	UserSignUp(receiver string, activationToken string) error
}
