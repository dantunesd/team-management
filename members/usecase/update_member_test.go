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

func TestUpdateMember_Handle(t *testing.T) {
	type fields struct {
		repository domain.MembersRepository
	}
	type args struct {
		id    string
		input *usecase.UpdateMemberInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *usecase.UpdateMemberOutput
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
				input: &usecase.UpdateMemberInput{
					Type:     "invalid",
					TypeData: []byte(``),
				},
			},
			want:    &usecase.UpdateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return an error if Validate fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					return mocks.NewMembersRepository(t)
				}(),
			},
			args: args{
				input: &usecase.UpdateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"unknown": "unknown"}`),
				},
			},
			want:    &usecase.UpdateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return an error if Update fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Update", mock.Anything).Return(false, errors.New("failed"))
					return repository
				}(),
			},
			args: args{
				input: &usecase.UpdateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"contract_duration": 10}`),
				},
			},
			want:    &usecase.UpdateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return an error if no member is updated",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Update", mock.Anything).Return(false, nil)
					return repository
				}(),
			},
			args: args{
				input: &usecase.UpdateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"contract_duration": 10}`),
				},
			},
			want:    &usecase.UpdateMemberOutput{},
			wantErr: true,
		},
		{
			name: "should update successfully",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Update", mock.Anything).Return(true, nil)
					return repository
				}(),
			},
			args: args{
				input: &usecase.UpdateMemberInput{
					Type:     "contractor",
					TypeData: []byte(`{"contract_duration": 10}`),
				},
			},
			want: &usecase.UpdateMemberOutput{
				UpdatedAt: utils.GetTimeNowInUTC(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := usecase.NewUpdateMember(tt.fields.repository)
			got, err := g.Handle(tt.args.id, tt.args.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMember.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMember.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
