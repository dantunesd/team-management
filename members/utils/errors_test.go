package utils

import (
	"reflect"
	"testing"
)

func TestCustomError_Error(t *testing.T) {
	type fields struct {
		Message string
		Code    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return the message",
			fields: fields{
				Message: "my customized message",
				Code:    400,
			},
			want: "my customized message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewError(tt.fields.Message, tt.fields.Code)
			if got := d.Error(); got != tt.want {
				t.Errorf("CustomError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomError_Code(t *testing.T) {
	type fields struct {
		message string
		code    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "should return the code",
			fields: fields{
				message: "message",
				code:    500,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CustomError{
				message: tt.fields.message,
				code:    tt.fields.code,
			}
			if got := d.Code(); got != tt.want {
				t.Errorf("CustomError.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBadRequest(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *CustomError
	}{
		{
			"should return a bad request custom error",
			args{
				message: "error",
			},
			&CustomError{
				message: "error",
				code:    400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBadRequest(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNotFound(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *CustomError
	}{
		{
			"should return a not found custom error",
			args{
				message: "error",
			},
			&CustomError{
				message: "error",
				code:    404,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotFound(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
