/* Copyright (C) Fedir Petryk */

package user

func (p Publisher) UserDeactivated(email string) {
	msg := Deactivated{
		Receiver: email,
	}

	p.publish(deactivated, msg)
}
