# merr
`merr` is library for handle multiple error on golang.

This allows a function ig Go to return an list of error.
It will be usefull when you do call some function to get multiple error handler. 

`merr` is fully compatible with the Go standard library
[errors](https://golang.org/pkg/errors/) package, including the
functions `As`, `Is`, and `Unwrap`. This provides a standardized approach
for introspecting on error values.

## Documentation
[![Go Doc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/firdasafridi/merr)
[![Go Report Card](https://goreportcard.com/badge/github.com/firdasafridi/merr)](https://goreportcard.com/report/github.com/firdasafridi/merr)
[![Build Status](http://img.shields.io/travis/firdasafridi/merr.svg?style=flat-square)](https://travis-ci.org/firdasafridi/merr)



## Installation

Install using `go get github.com/firdasafridi/merr`.


## Example
`merr` is easy to use in existing Go applications/libraries that may not be aware of it.

The `Set` function is used to create a list of errors. If you need to set prefix you can add it with `SetPrefix`.

```go
var mulerr merr.Error

if err := sampleLogic1(); err != nil {
	mulerr.Set(err)
}

if err := sampleLogic2(); err != nil {
	mulerr.SetPrefix("Sample Prefix", err)
}
    
return mulerr.IsError()
```


### Customizing the formating of the errors

If you want customizing format error it will be simple like this.

```go
var funcFormat = func(errorList []error) string {
 	if len(errorList) == 1 {
		return ""
	}

	var listErr string
	for _, err := range errorList {
		listErr += err.Error() + "\n"
	}
	return "[ERROR Change]" + listErr
}

merr.NewFormat(funcFormat)
```
