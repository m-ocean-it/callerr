package werror

import (
	"fmt"
)

// New creates a new error with the supplied message and annotates it with the caller's name.
func New(msg string, args ...any) error {
	callerName := getCaller()

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return fmt.Errorf("%s: %s", callerName, msg)
}

// Wrap annotates the error with the caller's name.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	callerName := getCaller()

	return fmt.Errorf("%s: %w", callerName, err)
}

// WrapWithMsg annotates the error with the caller's name and the supplied message.
func WrapWithMsg(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}

	callerName := getCaller()

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return fmt.Errorf("%s: %s: %w", callerName, msg, err)
}

// WithMsg annotates the error with the supplied message without the caller's name.
func WithMsg(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return fmt.Errorf("%s: %w", msg, err)
}
