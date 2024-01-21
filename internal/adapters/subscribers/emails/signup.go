/* Copyright (C) Fedir Petryk */

package user

import (
	"encoding/json"
)

func (s Subscriber) Signup(data []byte) {
	userDeact := &SignUp{}
	err := json.Unmarshal(data, userDeact)
	if err != nil {
		s.logger.Errorw("error unmarshalling user deactivated data", "error", err)
		return
	}

	err = s.service.UserSignUp(userDeact.Receiver, userDeact.ActivationToken)
	if err != nil {
		s.logger.Errorw("error sending user deactivated email", "error", err)
	}
}
