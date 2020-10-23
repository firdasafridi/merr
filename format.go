package merr

// FormatErr is a basic formatter that outputs the error.
func FormatErr(errorList []error) string {
	if len(errorList) == 0 {
		return ""
	}

	var listErr string
	for _, err := range errorList {
		listErr += err.Error() + "\n"
	}

	return listErr
}

// Len error from merr.Errors
func Len(err error) int {
	if err == nil {
		return 0
	}
	switch err := err.(type) {
	case *Error:
		return len(err.Errors)
	}
	return 1
}

// Check will return multi error format if the data is true
func Check(err error) *Error {
	if err == nil {
		return nil
	}
	switch err := err.(type) {
	case *Error:
		return err
	}
	return nil
}
