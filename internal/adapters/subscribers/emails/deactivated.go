/* Copyright (C) Fedir Petryk */

package user

import (
	"encoding/json"
)

func (s Subscriber) Deactivated(data []byte) {
	userDeact := &Deactivated{}
	err := json.Unmarshal(data, userDeact)
	if err != nil {
		s.logger.Errorw("error unmarshalling user deactivated data", "error", err)
		return
	}

	err = s.service.UserDeactivated(userDeact.Receiver)
	if err != nil {
		s.logger.Errorw("error sending user deactivated email", "error", err)
	}
}
