package merr

import (
	"errors"
	"reflect"
	"testing"
)

func TestFormatErr(t *testing.T) {
	type args struct {
		errorList []error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should be return list of error string",
			args: args{
				errorList: []error{
					errors.New("errFoo1"),
					errors.New("errFoo2"),
				},
			},
			want: "errFoo1\nerrFoo2\n",
		},
		{
			name: "Should be return empty string",
			args: args{
				errorList: []error{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatErr(tt.args.errorList); got != tt.want {
				t.Errorf("FormatErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should be return error common",
			args: args{
				err: errors.New("errFoo"),
			},
			want: 1,
		},
		{
			name: "Should be return 2 error from format",
			args: args{
				err: &Error{
					Errors: []error{
						errors.New("errFoo1"),
						errors.New("errFoo2"),
					},
				},
			},
			want: 2,
		},
		{
			name: "Should be return 0 because no error",
			args: args{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Len(tt.args.err); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "Should be return *Error",
			args: args{
				err: &Error{},
			},
			want: &Error{},
		},
		{
			name: "Should be return nil because error is not error format",
			args: args{
				err: errors.New("errFoo"),
			},
			want: nil,
		},
		{
			name: "Should be return nil because no error",
			args: args{
				err: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
