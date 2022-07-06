package usecase_test

import (
	"errors"
	"reflect"
	"team-management/members/domain"
	"team-management/members/usecase"
	"team-management/mocks"
	"testing"
)

func TestGetMember_Handle(t *testing.T) {
	type fields struct {
		repository domain.MembersRepository
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *usecase.GetMemberOutput
		wantErr bool
	}{
		{
			name: "should return an error if Get fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Get", "1").Return(&domain.Member{}, errors.New("failed"))
					return repository
				}(),
			},
			args: args{
				ID: "1",
			},
			want:    &usecase.GetMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return an error if no member is found",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Get", "1").Return(&domain.Member{}, nil)
					return repository
				}(),
			},
			args: args{
				ID: "1",
			},
			want:    &usecase.GetMemberOutput{},
			wantErr: true,
		},
		{
			name: "should return a member correctly",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Get", "1").Return(&domain.Member{
						ID:        "1",
						Name:      "whatever",
						Type:      domain.EmployeeType,
						TypeData:  &domain.Employee{Role: "se"},
						Tags:      domain.Tags{"backend", "frontend"},
						CreatedAt: "now",
						UpdatedAt: "tomorrow",
					}, nil)
					return repository
				}(),
			},
			args: args{
				ID: "1",
			},
			want: &usecase.GetMemberOutput{
				ID:        "1",
				Name:      "whatever",
				Type:      "employee",
				TypeData:  []byte(`{"role":"se"}`),
				Tags:      []string{"backend", "frontend"},
				CreatedAt: "now",
				UpdatedAt: "tomorrow",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := usecase.NewGetMember(tt.fields.repository)
			got, err := g.Handle(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMember.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMember.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
