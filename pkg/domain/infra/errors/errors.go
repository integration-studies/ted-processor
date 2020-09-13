package errors

import (
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func Wrap(err error) error {
	if _, ok := err.(stackTracer); ok {
		return err
	}
	return errors.WithStack(err)
}

func WrapWithMessage(err error, msg string) error {
	if _, ok := err.(stackTracer); ok {
		return errors.WithMessage(err, msg)
	}
	return errors.Wrap(err, msg)
}

func WrapWithMessagef(err error, msg string, args ...interface{}) error {
	if _, ok := err.(stackTracer); ok {
		return errors.WithMessagef(err, msg, args...)
	}
	return errors.Wrapf(err, msg, args...)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func New(msg string) error {
	return errors.New(msg)
}
