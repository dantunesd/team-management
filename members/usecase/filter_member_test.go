package usecase_test

import (
	"errors"
	"reflect"
	"team-management/members/domain"
	"team-management/members/usecase"
	"team-management/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestFilterMember_Handle(t *testing.T) {
	type fields struct {
		repository domain.MembersRepository
	}
	type args struct {
		input *usecase.FilterMemberInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []usecase.FilterMemberOutput
		wantErr bool
	}{
		{
			name: "should return an error if Filter fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Filter", mock.Anything).Return([]*domain.Member{}, errors.New("failed"))
					return repository
				}(),
			},
			args: args{
				input: &usecase.FilterMemberInput{},
			},
			want:    []usecase.FilterMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return members successfully",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Filter", mock.Anything).Return([]*domain.Member{
						{ID: "1", Name: "whatever", Type: domain.EmployeeType, TypeData: &domain.Employee{Role: "se"}, Tags: domain.Tags{"backend"}, CreatedAt: "now", UpdatedAt: "tomorrow"},
						{ID: "2", Name: "whatever", Type: domain.ContractorType, TypeData: &domain.Contractor{ContractDuration: 10}, Tags: domain.Tags{"backend"}, CreatedAt: "now", UpdatedAt: "tomorrow"},
					}, nil)
					return repository
				}(),
			},
			args: args{
				input: &usecase.FilterMemberInput{},
			},
			want: []usecase.FilterMemberOutput{
				{ID: "1", Name: "whatever", Type: "employee", TypeData: []byte(`{"role":"se"}`), Tags: domain.Tags{"backend"}, CreatedAt: "now", UpdatedAt: "tomorrow"},
				{ID: "2", Name: "whatever", Type: "contractor", TypeData: []byte(`{"contract_duration":10}`), Tags: domain.Tags{"backend"}, CreatedAt: "now", UpdatedAt: "tomorrow"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := usecase.NewFilterMember(tt.fields.repository)
			got, err := f.Handle(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterMember.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMember.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
