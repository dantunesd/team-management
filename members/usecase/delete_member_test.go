package usecase_test

import (
	"errors"
	"team-management/members/domain"
	"team-management/members/usecase"
	"team-management/mocks"
	"testing"
)

func TestDeleteMember_Handle(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "should return an error if Delete fails",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Delete", "1").Return(false, errors.New("failed"))
					return repository
				}(),
			},
			args: args{
				ID: "1",
			},
			wantErr: true,
		},
		{
			name: "should return an error if no member is deleted",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Delete", "1").Return(false, nil)
					return repository
				}(),
			},
			args: args{
				ID: "1",
			},
			wantErr: true,
		},
		{
			name: "should delete a member successfully",
			fields: fields{
				repository: func() domain.MembersRepository {
					repository := mocks.NewMembersRepository(t)
					repository.On("Delete", "1").Return(true, nil)
					return repository
				}(),
			},
			args: args{
				ID: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := usecase.NewDeleteMember(tt.fields.repository)
			if err := g.Handle(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMember.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
