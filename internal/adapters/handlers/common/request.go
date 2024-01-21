/* Copyright (C) Fedir Petryk */

package common

type BaseSearchRequest struct {
	Page int `json:"page" validate:"number,max=100000" json:"page,omitempty"`
	Size int `json:"size" validate:"number,max=1000" json:"size,omitempty"`
	Sort int `json:"sort" validate:"omitempty,number,min=1,max=15" json:"size,omitempty"`
} // @name BaseSearchRequest
