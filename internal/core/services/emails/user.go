/* Copyright (C) Fedir Petryk */

package emails

import (
	"fmt"

	"go.uber.org/zap"
)

type User struct {
	emailer Emailer
	logger  *zap.SugaredLogger
}

func NewUserEmailsSrv(emailer Emailer, logger *zap.SugaredLogger) User {
	return User{
		emailer: emailer,
		logger:  logger,
	}
}

func (u User) UserDeactivated(receiver string) error {
	err := u.emailer.Send("User Deactivated", receiver, fmt.Sprintf(`Profile deactivated`))
	if err != nil {
		u.logger.Errorw(err.Error(), "method", "SignUp")
	}

	return nil
}

func (u User) UserSignUp(receiver string, activationToken string) error {
	err := u.emailer.Send("User SignUp", receiver, fmt.Sprintf(`Activation link: %s`, activationToken))
	if err != nil {
		u.logger.Errorw(err.Error(), "method", "SignUp")
	}

	return nil
}
