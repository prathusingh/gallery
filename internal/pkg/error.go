package pkg

import (
	"fmt"
)

type ErrorCode uint

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeInvalidArgument
)

type Error struct {
	orig error
	msg  string
	code ErrorCode
}

// WrapError returns a wrapped error
func WrapError(orig error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		code: code,
		orig: orig,
		msg:  fmt.Sprintf(format, a...),
	}
}

// NewError instantiates a new error
func NewError(code ErrorCode, format string, a ...interface{}) error {
	return WrapError(nil, code, format, a...)
}

// Error returns the message, when wrapping errors the wrapping error is returned
func (e *Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

// Unwrap returns the wrapped error, if any
func (e *Error) Unwrap() error {
	return e.orig
}

// Code represents the code representing this error.
func (e *Error) Code() ErrorCode {
	return e.code
}
