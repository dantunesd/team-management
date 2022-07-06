package usecase_test

import (
	"errors"
	"reflect"
	"team-management/members/domain"
	"team-management/members/usecase"
	"team-management/members/utils"
	"team-management/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCreateMember_Handle(t *testing.T) {
	type fields struct {
		repository domain.MembersRepository
	}
	type args struct {
		input *usecase.CreateMemberInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *usecase.CreateMemberOutput
		wantErr bool
	}{
		{
			name: "should return an error if TypeDataFactory fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					return mocks.NewMembersRepository(t)
				}(),
			},
			args: args{
				input: &usecase.CreateMemberInput{
					Type:     "invalid",
					TypeData: []byte(``),
				},
			},
			want:    &usecase.CreateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return an error if NewMember fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					return mocks.NewMembersRepository(t)
				}(),
			},
			args: args{
				input: &usecase.CreateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"unknown": "unknown"}`),
				},
			},
			want:    &usecase.CreateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return an error Create fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Create", mock.Anything).Return("", errors.New("failed"))
					return repository
				}(),
			},
			args: args{
				input: &usecase.CreateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"contract_duration": 10}`),
				},
			},
			want:    &usecase.CreateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should create a member successfully",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Create", mock.Anything).Return("id-generated", nil)
					return repository
				}(),
			},
			args: args{
				input: &usecase.CreateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"contract_duration": 10}`),
				},
			},
			want: &usecase.CreateMemberOutput{
				ID:        "id-generated",
				CreatedAt: utils.GetTimeNowInUTC(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := usecase.NewCreateMember(tt.fields.repository)
			got, err := c.Handle(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMember.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMember.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
