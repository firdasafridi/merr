// Package merr is library for handle multiple error on golang
// (see https://pkg.go.dev/github.com/firdasafridi/merr?tab=overview) for more documentation
//
// This allows a function ig Go to return an list of error.
// It will be useful when you do call some function to get multiple error handler.
//
//
// merr is fully compatible with the Go standard library errors package, including the functions As, Is, and Unwrap.
// This provides a standardized approach for introspecting on error values.
package merr
