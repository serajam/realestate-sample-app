/* Copyright (C) Fedir Petryk */

package emails

type Emailer interface {
	Send(subject, receiver string, body string) error
}
