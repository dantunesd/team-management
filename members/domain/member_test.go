package domain

import (
	"team-management/members/utils"
	"testing"
)

func TestMember_Validate(t *testing.T) {
	type args struct {
		name      string
		typeName  Type
		typeData  TypeData
		tags      Tags
		createdAt string
		updatedAt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return nil if it's a Employee member",
			args: args{
				name:      "example",
				typeName:  EmployeeType,
				typeData:  &Employee{Role: "Software Eng"},
				tags:      Tags{},
				createdAt: utils.GetTimeNowInUTC(),
				updatedAt: utils.GetTimeNowInUTC(),
			},
			wantErr: false,
		},
		{
			name: "should return nil if it's a Contractor member",
			args: args{
				name:      "example",
				typeName:  ContractorType,
				typeData:  &Contractor{ContractDuration: 1},
				tags:      Tags{},
				createdAt: utils.GetTimeNowInUTC(),
				updatedAt: utils.GetTimeNowInUTC(),
			},
			wantErr: false,
		},
		{
			name: "should return an error if type is invalid",
			args: args{
				name:     "example",
				typeName: "whatever",
				typeData: &Employee{Role: "Software Eng"},
				tags:     Tags{},
			},
			wantErr: true,
		},
		{
			name: "should return an error if typeData is invalid",
			args: args{
				name:      "example",
				typeName:  EmployeeType,
				typeData:  &Employee{},
				tags:      Tags{},
				createdAt: utils.GetTimeNowInUTC(),
				updatedAt: utils.GetTimeNowInUTC(),
			},
			wantErr: true,
		},
		{
			name: "should return an error if type is different of typeData",
			args: args{
				name:      "example",
				typeName:  ContractorType,
				typeData:  &Employee{},
				tags:      Tags{},
				createdAt: utils.GetTimeNowInUTC(),
				updatedAt: utils.GetTimeNowInUTC(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Member{
				Name:      tt.name,
				Type:      tt.args.typeName,
				TypeData:  tt.args.typeData,
				Tags:      tt.args.tags,
				CreatedAt: tt.args.createdAt,
				UpdatedAt: tt.args.updatedAt,
			}

			err := got.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
