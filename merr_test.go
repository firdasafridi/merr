package merr

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

var errFoo = errors.New("errFoo")

func TestNewFormat(t *testing.T) {
	type args struct {
		fErr func(errorList []error) string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should be change format",
			args: args{
				fErr: func(errorList []error) string {
					return ""
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewFormat(tt.args.fErr)
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		Errors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should be  return format string",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Errors: tt.fields.Errors,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Set(t *testing.T) {
	type fields struct {
		Errors []error
	}
	type args struct {
		newErr []error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Should be set new error",
			args: args{
				newErr: []error{},
			},
		},
		{
			name: "Should be not set new error because no error",
			args: args{
				newErr: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Errors: tt.fields.Errors,
			}
			err.Set(tt.args.newErr...)
		})
	}
}

func TestError_SetPrefix(t *testing.T) {
	type fields struct {
		Errors []error
	}
	type args struct {
		prefix string
		newErr error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Should be set error with prefix",
			args: args{
				prefix: "ERR",
				newErr: errFoo,
			},
		},
		{
			name: "Should be not set error because error nil",
			args: args{
				prefix: "ERR",
				newErr: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Errors: tt.fields.Errors,
			}
			err.SetPrefix(tt.args.prefix, tt.args.newErr)
		})
	}
}

func TestError_Len(t *testing.T) {
	type fields struct {
		Errors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Should be get 0 error",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Errors: tt.fields.Errors,
			}
			if got := err.Len(); got != tt.want {
				t.Errorf("Error.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_IsError(t *testing.T) {
	tests := []struct {
		name    string
		fields  *Error
		wantErr bool
	}{
		{
			name: "Should be get errors",
			fields: &Error{
				Errors: []error{
					errFoo,
				},
			},
			wantErr: true,
		},
		{
			name: "Should be no errors",
			fields: &Error{
				Errors: []error{},
			},
			wantErr: false,
		},
		{
			name:    "Should be no errors",
			fields:  nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields
			if err := err.IsError(); (err != nil) != tt.wantErr {
				t.Errorf("Error.IsError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestError_WrappedErrors(t *testing.T) {
	type fields struct {
		Errors []error
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		{
			name: "Should be get wrapping errors",
			fields: fields{
				Errors: []error{
					errFoo,
				},
			},
			want: []error{
				errFoo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Errors: tt.fields.Errors,
			}
			if got := err.WrappedErrors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error.WrappedErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	type fields struct {
		Errors []error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Should be return unwraping error",
			fields: fields{
				Errors: []error{
					errFoo,
					errFoo,
				},
			},
			wantErr: true,
		},
		{
			name: "Should be return unwraping error",
			fields: fields{
				Errors: []error{
					errFoo,
				},
			},
			wantErr: true,
		},
		{
			name: "Should be return unwraping no error",
			fields: fields{
				Errors: []error{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Errors: tt.fields.Errors,
			}
			if err := err.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("Error.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_chain_Error(t *testing.T) {
	tests := []struct {
		name string
		e    chain
		want string
	}{
		{
			name: "Should be return error",
			e: chain{
				errFoo,
			},
			want: "errFoo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("chain.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chain_Unwrap(t *testing.T) {
	tests := []struct {
		name    string
		e       chain
		wantErr bool
	}{
		{
			name: "Should be unwrap error chain",
			e: chain{
				errFoo,
				errFoo,
			},
			wantErr: true,
		},
		{
			name: "Should be unwrap error chain",
			e: chain{
				errFoo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("chain.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
