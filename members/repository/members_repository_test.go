package repository

import (
	"errors"
	"reflect"
	"team-management/members/domain"
	"team-management/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMembersRepository_Create(t *testing.T) {
	type fields struct {
		db Database
	}
	type args struct {
		member *domain.Member
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return an error if db.Create fails",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Create", mock.Anything).Return("", errors.New("failed"))
					return db
				}(),
			},
			args: args{
				member: &domain.Member{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should create successfully",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Create", mock.Anything).Return("id-generated", nil)
					return db
				}(),
			},
			args: args{
				member: &domain.Member{},
			},
			want:    "id-generated",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMembersRepository(tt.fields.db)
			got, err := m.Create(tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MembersRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembersRepository_Get(t *testing.T) {
	type fields struct {
		db Database
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Member
		wantErr bool
	}{
		{
			name: "should return an error if db.Get fails",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Get", idField, "1", &MemberDB{}).Return(errors.New("failed"))
					return db
				}(),
			},
			args: args{
				ID: "1",
			},
			want:    &domain.Member{},
			wantErr: true,
		},
		{
			name: "should return a domain member filled correctly",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Get", idField, "1", &MemberDB{}).Return(nil).Run(func(args mock.Arguments) {
						memberDB := args.Get(2).(*MemberDB)
						memberDB.ID = "1"
						memberDB.Name = "name"
					})
					return db
				}(),
			},
			args: args{
				ID: "1",
			},
			want: &domain.Member{
				ID:   "1",
				Name: "name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMembersRepository(tt.fields.db)
			got, err := m.Get(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MembersRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembersRepository_Delete(t *testing.T) {
	type fields struct {
		db Database
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "should return an error if db.Delete fails",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Delete", idField, "1").Return(false, errors.New("failed"))
					return db
				}(),
			},
			args: args{
				ID: "1",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should return false if any item wasn't deleted",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Delete", idField, "1").Return(false, nil)
					return db
				}(),
			},
			args: args{
				ID: "1",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "should return true if any item was deleted",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Delete", idField, "1").Return(true, nil)
					return db
				}(),
			},
			args: args{
				ID: "1",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMembersRepository(tt.fields.db)
			got, err := m.Delete(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MembersRepository.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembersRepository_Update(t *testing.T) {
	type fields struct {
		db Database
	}
	type args struct {
		member *domain.Member
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "should return an error if db.Update fails",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Update", idField, "1", mock.Anything).Return(false, errors.New("failed"))
					return db
				}(),
			},
			args: args{
				member: &domain.Member{ID: "1"},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "should update successfully",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Update", idField, "1", mock.Anything).Return(true, nil)
					return db
				}(),
			},
			args: args{
				member: &domain.Member{ID: "1"},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MembersRepository{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MembersRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembersRepository_Filter(t *testing.T) {
	type fields struct {
		db Database
	}
	type args struct {
		filters map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Member
		wantErr bool
	}{
		{
			name: "should return an error if db.Filter fails",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Filter", mock.Anything, mock.Anything).Return(errors.New("failed"))
					return db
				}(),
			},
			args: args{
				filters: map[string]string{"name": "example"},
			},
			want:    []*domain.Member{},
			wantErr: true,
		},
		{
			name: "should return an empty list if any member wasn't found",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Filter", mock.Anything, mock.Anything).Return(nil)
					return db
				}(),
			},
			args: args{
				filters: map[string]string{"name": "example"},
			},
			want:    []*domain.Member{},
			wantErr: false,
		},
		{
			name: "should return an error if filters contain invalid fields",
			fields: fields{
				db: func() Database {
					return mocks.NewDatabase(t)
				}(),
			},
			args: args{
				filters: map[string]string{"invalid": "value"},
			},
			want:    []*domain.Member{},
			wantErr: true,
		},
		{
			name: "should return a filled list if many members was found",
			fields: fields{
				db: func() Database {
					db := mocks.NewDatabase(t)
					db.On("Filter", mock.Anything, &[]*MemberDB{}).Return(nil).Run(func(args mock.Arguments) {
						membersDB := args.Get(1).(*[]*MemberDB)
						*membersDB = append(*membersDB, &MemberDB{ID: "1"})
						*membersDB = append(*membersDB, &MemberDB{ID: "1"})
					})
					return db
				}(),
			},
			args: args{
				filters: map[string]string{"name": "example"},
			},
			want: []*domain.Member{
				NewMemberDomain(&MemberDB{ID: "1"}),
				NewMemberDomain(&MemberDB{ID: "1"}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MembersRepository{
				db: tt.fields.db,
			}
			got, err := m.Filter(tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersRepository.Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MembersRepository.Filter() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
