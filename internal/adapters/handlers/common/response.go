/* Copyright (C) Fedir Petryk */

package common

import (
	"github.com/go-playground/validator/v10"
)

type EmptyResponse struct {
} // @name EmptyResponse

type DefaultResponse struct {
	Count      int         `json:"count,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	Data       interface{} `json:"data"`
} // @name DefaultResponseContainer

type ErrorResponse struct {
	Message string `json:"error,omitempty"`
} // @name ErrorResponse

type ErrorValidateResponse struct {
	Err         string       `json:"error"`
	FieldErrors []FieldError `json:"fieldErrors"`
} // @name ValidationErrorResponse

func NewErrorValidateResponse(message string, fieldErrors validator.ValidationErrors) ErrorValidateResponse {
	err := ErrorValidateResponse{
		Err: message,
	}

	for _, fieldError := range fieldErrors {
		err.FieldErrors = append(
			err.FieldErrors, FieldError{
				Error: fieldError.Error(),
				Field: fieldError.Field(),
				Tag:   fieldError.Tag(),
			},
		)
	}

	return err
}

type FieldError struct {
	Error string `json:"error"`
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

type ModelDateTime struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}
