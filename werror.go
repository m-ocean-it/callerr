package werror

import (
	"errors"
)

// New creates a new error with the supplied message and annotates it with the caller's name.
func New(msg string) error {
	callerName := getCaller()

	return errors.New(callerName + ": " + msg)
}

// Wrap annotates the error with the caller's name.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	callerName := getCaller()

	return &wrapError{
		msg: callerName + ": " + err.Error(),
		err: err,
	}
}

// WrapWithMsg annotates the error with the caller's name and the supplied message.
func WrapWithMsg(err error, msg string) error {
	if err == nil {
		return nil
	}

	callerName := getCaller()

	return &wrapError{
		msg: callerName + ": " + msg + ": " + err.Error(),
		err: err,
	}
}

// WithMsg annotates the error with the supplied message without the caller's name.
func WithMsg(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &wrapError{
		msg: msg + ": " + err.Error(),
		err: err,
	}
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}
