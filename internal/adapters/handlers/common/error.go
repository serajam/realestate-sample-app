/* Copyright (C) Fedir Petryk */

package common

import (
	"errors"
	"fmt"
	"net/http"
)

// Error codes
// Only common use codes should be in the library, implementation-specific errors should be in other packages.
var (
	ErrNotFound        = New("NOT_FOUND", http.StatusNotFound)
	ErrExists          = New("EXIST_ERROR", http.StatusBadRequest)
	ErrInternal        = New("INTERNAL_ERROR", http.StatusInternalServerError)
	ErrTimeout         = New("REQUEST_TIMEOUT", http.StatusRequestTimeout)
	ErrUnauthorized    = New("UNAUTHORIZED", http.StatusUnauthorized)
	ErrBind            = New("COULD_NOT_PARSE_AND_BIND_REQUEST", http.StatusBadRequest)
	ErrMapping         = New("COULD_NOT_MAP_DATA", http.StatusUnprocessableEntity)
	ErrValidate        = New("REQUEST_VALIDATION_FAILURE", http.StatusBadRequest)
	ErrAccessForbidden = New("ACCESS_FORBIDDEN", http.StatusForbidden)
)

type Error struct {
	Err  string `json:"error"`
	Code int    `json:"-"`

	// chain error
	err error
}

func New(message string, code int) Error {
	return Error{
		Err:  message,
		Code: code,
	}
}

// Error implements error interface
func (e Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.Err, e.err)
	}

	return e.Err
}

// Wrap saves error to the chain.
func (e Error) Wrap(err error) Error {
	e.err = err

	return e
}

// Unwrap returns the next error in the error chain.
func (e Error) Unwrap() error {
	return e.err
}

// Is reports whether any error in the chain matches target.
func (e Error) Is(err error) bool {
	var codeErr Error
	if errors.As(err, &codeErr) {
		return codeErr.Code == e.Code && codeErr.Err == e.Err
	}

	return false
}
