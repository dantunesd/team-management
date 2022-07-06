package domain

import (
	"testing"
)

func TestEmployee_IsValid(t *testing.T) {
	type fields struct {
		Role string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return true",
			fields: fields{
				Role: "Software Engineer",
			},
			want: true,
		},
		{
			name: "should return false",
			fields: fields{
				Role: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{
				Role: tt.fields.Role,
			}
			if got := e.IsValid(); got != tt.want {
				t.Errorf("Employee.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_GetType(t *testing.T) {
	tests := []struct {
		name string
		want Type
	}{
		{
			name: "should return Employee",
			want: "employee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{}
			if got := e.GetType(); got != tt.want {
				t.Errorf("Employee.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
