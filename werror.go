package err

import (
	"fmt"
)

func New(msg string, args ...any) error {
	callerName := getCaller()

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return fmt.Errorf("%s: %s", callerName, msg)
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	callerName := getCaller()

	return fmt.Errorf("%s: %w", callerName, err)
}

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
