/* Copyright (C) Fedir Petryk */

package errors

const (
	msgTimeout      = "REQUEST_TIMEOUT"
	msgExists       = "ENTIRY_EXISTS"
	msgNotFound     = "NOT_FOUND"
	msgAccessDenied = "ACCESS_DENIED"
	msgInternal     = "INTERNAL_ERROR"

	MsgFailCreateOp = "FAIL_CREATE"
	MsgFailGetOp    = "FAIL_GET"
	MsgFailUpdateOp = "FAIL_UPDATE"
	MsgFailDeleteOp = "FAIL_DELETE"
)

const (
	MsgUserSingupFailed = "SIGN_UP_FAIL"
	MsgUserNotExists    = "USER_NOT_EXISTS"
	MsgEmailExists      = "EMAIL_EXISTS"
	MsgWrongEmailOrPwd  = "WRONG_EMAIL_OR_PASSWORD"
	MsgInvalidToken     = "INVALID_TOKEN"
)
