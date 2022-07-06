package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTypeDataFactory(t *testing.T) {
	type args struct {
		typeName Type
		data     json.RawMessage
	}
	tests := []struct {
		name    string
		args    args
		want    TypeData
		wantErr bool
	}{
		{
			name: "should return a contractor",
			args: args{
				typeName: ContractorType,
				data:     []byte(`{"contract_duration": 10}`),
			},
			want: &Contractor{ContractDuration: 10},

			wantErr: false,
		},
		{
			name: "should return an employee",
			args: args{
				typeName: EmployeeType,
				data:     []byte(`{"role": "Software Engineer"}`),
			},
			want:    &Employee{Role: "Software Engineer"},
			wantErr: false,
		},
		{
			name: "should return an error if type is invalid",
			args: args{
				typeName: "invalid",
				data:     json.RawMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return an error if type data is invalid",
			args: args{
				typeName: EmployeeType,
				data:     nil,
			},
			want:    &Employee{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TypeDataFactory(tt.args.typeName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeDataFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeDataFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
