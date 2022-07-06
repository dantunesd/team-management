package domain

import (
	"testing"
)

func TestContractor_IsValid(t *testing.T) {
	type fields struct {
		ContractDuration int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return true",
			fields: fields{
				ContractDuration: 1,
			},
			want: true,
		},
		{
			name: "should return false",
			fields: fields{
				ContractDuration: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contractor{
				ContractDuration: tt.fields.ContractDuration,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("Contractor.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContractor_GetType(t *testing.T) {
	tests := []struct {
		name string
		want Type
	}{
		{
			name: "should return Contractor",
			want: "contractor",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contractor{}
			if got := c.GetType(); got != tt.want {
				t.Errorf("Contractor.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
