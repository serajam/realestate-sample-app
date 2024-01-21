/* Copyright (C) Fedir Petryk */

package properties

type UserSavedHome struct {
	PropertyID int `bun:"property_id"`
	UserID     int `bun:"user_id"`
}
