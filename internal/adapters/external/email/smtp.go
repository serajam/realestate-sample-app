/* Copyright (C) Fedir Petryk */

package email

import (
	"net/smtp"

	"github.com/labstack/gommon/email"
	"github.com/pkg/errors"
)

type SMTP struct {
	sender  string
	emailer *email.Email
}

func NewSMTP(host string, port string, login string, pwd string, sender string) SMTP {
	emailer := email.New(host + ":" + port)
	emailer.Auth = smtp.PlainAuth("", login, pwd, host)
	return SMTP{
		sender:  sender,
		emailer: emailer,
	}
}

func (s SMTP) Send(subject, receiver string, body string) error {
	m := email.Message{
		From:     s.sender,
		To:       receiver,
		Subject:  subject,
		BodyText: body,
	}
	err := s.emailer.Send(&m)
	if err != nil {
		return errors.Wrap(err, "smtp send")
	}

	return nil
}
