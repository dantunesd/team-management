package repository

import (
	"encoding/json"
	"reflect"
	"team-management/members/domain"
	"team-management/members/utils"
	"testing"
)

func TestNewMemberDB(t *testing.T) {
	convertedType, _ := json.Marshal(&domain.Contractor{})

	type args struct {
		member *domain.Member
	}
	tests := []struct {
		name string
		args args
		want *MemberDB
	}{
		{
			name: "should return a converted memberDB",
			args: args{
				member: &domain.Member{
					Name:      "whatever",
					Type:      domain.ContractorType,
					Tags:      make([]string, 0),
					TypeData:  &domain.Contractor{},
					CreatedAt: utils.GetTimeNowInUTC(),
					UpdatedAt: utils.GetTimeNowInUTC(),
				},
			},
			want: &MemberDB{
				Name:      "whatever",
				Type:      "contractor",
				Tags:      make([]string, 0),
				TypeData:  convertedType,
				CreatedAt: utils.GetTimeNowInUTC(),
				UpdatedAt: utils.GetTimeNowInUTC(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemberDB(tt.args.member); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemberDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMemberDomain(t *testing.T) {
	type args struct {
		member *MemberDB
	}
	tests := []struct {
		name string
		args args
		want *domain.Member
	}{
		{
			name: "should return a converted domain.Member",
			args: args{
				member: &MemberDB{
					Name:      "whatever",
					Type:      "employee",
					Tags:      make([]string, 0),
					TypeData:  []byte(`{"role":"se"}`),
					CreatedAt: utils.GetTimeNowInUTC(),
					UpdatedAt: utils.GetTimeNowInUTC(),
				},
			},
			want: &domain.Member{
				Name:      "whatever",
				Type:      domain.EmployeeType,
				Tags:      make([]string, 0),
				TypeData:  &domain.Employee{Role: "se"},
				CreatedAt: utils.GetTimeNowInUTC(),
				UpdatedAt: utils.GetTimeNowInUTC(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemberDomain(tt.args.member); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemberDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
