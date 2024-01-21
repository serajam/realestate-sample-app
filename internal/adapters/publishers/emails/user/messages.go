/* Copyright (C) Fedir Petryk */

package user

type Deactivated struct {
	Receiver string
}

type SignUp struct {
	Receiver        string
	ActivationToken string
}
