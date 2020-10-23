package merr

import (
	"fmt"

	"github.com/pkg/errors"
)

var formatErr = FormatErr

// Error is containt field to tract multiple errors. This accumulate errors.
type Error struct {
	Errors []error
}

// NewFormat change default format error it used
func NewFormat(fErr func(errorList []error) string) {
	formatErr = fErr
}

func (err *Error) Error() string {
	return formatErr(err.Errors)
}

// Set is function to set new error if not found it will return no error
func (err *Error) Set(newErr ...error) {
	if newErr == nil {
		return
	}
	err.Errors = append(err.Errors, newErr...)
}

// SetPrefix is function to set new error with prefix if not found it will return no error
func (err *Error) SetPrefix(prefix string, newErr error) {
	if newErr == nil {
		return
	}
	err.Errors = append(err.Errors, fmt.Errorf("[%s] %s", prefix, newErr))
}

// Len returns count errors represent list
func (err *Error) Len() int {
	return Len(err)
}

// IsError returns an error if errors represent a list of errors, or return nil if the list of error is empty.
// This function is useful at the end of the accumulation to handling the error.
func (err *Error) IsError() error {
	if err == nil {
		return nil
	}
	if len(err.Errors) == 0 {
		return nil
	}
	return err
}

// WrappedErrors returns the list of errors that this Error is wrapping.
// It is an implementation of the errwrap.Wrapper interface so that
// merr.Error can be used with that library.
func (err *Error) WrappedErrors() []error {
	return err.Errors
}

// Unwrap returns an error from Error (or nil if there are no errors).
// This error returned will further support Unwrap to get the next error,
// etc. The order will match the order of Errors in the merr.Error
// at the time of calling.
func (err *Error) Unwrap() error {
	// If we have no errors then we do nothing
	if err == nil || len(err.Errors) == 0 {
		return nil
	}

	// If we have exactly one error, we can just return that directly.
	if len(err.Errors) == 1 {
		return err.Errors[0]
	}

	// Shallow copy the slice
	errs := make([]error, len(err.Errors))
	copy(errs, err.Errors)
	return chain(errs)
}

// chain implements the interfaces necessary for errors.Is/As/Unwrap to
// work in a deterministic way with merr. A chain tracks a list of
// errors while accounting for the current represented error. This lets
// Is/As be meaningful.
type chain []error

// Error implements the error interface
func (e chain) Error() string {
	return e[0].Error()
}

// Unwrap implements errors.Unwrap by returning the next error in the
// chain or nil if there are no more errors.
func (e chain) Unwrap() error {
	if len(e) == 1 {
		return nil
	}

	return e[1:]
}

// As implements errors.As by attempting to map to the current value.
func (e chain) As(target interface{}) bool {
	return errors.As(e[0], target)
}

// Is implements errors.Is by comparing the current value directly.
func (e chain) Is(target error) bool {
	return errors.Is(e[0], target)
}
