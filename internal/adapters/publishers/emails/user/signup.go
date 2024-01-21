/* Copyright (C) Fedir Petryk */

package user

func (p Publisher) SignUp(receiver string, activationToken string) {
	msg := SignUp{
		Receiver:        receiver,
		ActivationToken: activationToken,
	}

	p.publish(signup, msg)
}
