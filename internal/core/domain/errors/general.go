/* Copyright (C) Fedir Petryk */

package errors

var (
	NotFound       = errNotFound{}
	Exists         = errEntityExists{}
	RequestTimeout = errRequestTimeout{}
	AccessDenied   = errAccessDenied{}
	Internal       = errInternal{}
	Dummy          = errDummy{}
)

// Dummy Return this error when you want to hide the fact that smth is happend but don't want to return nil
// don't forget to add actual error to the logs if it is important
type errDummy struct {
}

func (e errDummy) Error() string {
	return ""
}

type errInternal struct {
}

func (e errInternal) Error() string {
	return msgInternal
}

type errNotFound struct {
}

func (e errNotFound) Error() string {
	return msgNotFound
}

type errEntityExists struct {
}

func (e errEntityExists) Error() string {
	return msgExists
}

type errRequestTimeout struct {
}

func (e errRequestTimeout) Error() string {
	return msgTimeout
}

type errAccessDenied struct {
}

func (e errAccessDenied) Error() string {
	return msgAccessDenied
}

type OpFail struct {
	Op string
}

func (e OpFail) Error() string {
	return e.Op
}

type User struct {
	Msg string
}

func (e User) Error() string {
	return e.Msg
}
