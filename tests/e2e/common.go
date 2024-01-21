package e2e

import "encoding/json"

type ErrorResponse struct {
	FieldErrors []FieldError `json:"fieldErrors"`
	Message     string       `json:"message"`
	Error       string       `json:"error"`
}

type FieldError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Err   string `json:"error"`
}

type DefaultResponse struct {
	Count      int             `json:"count,omitempty"`
	TotalCount int             `json:"totalCount,omitempty"`
	Data       json.RawMessage `json:"data"`
}
