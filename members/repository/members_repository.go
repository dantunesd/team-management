package repository

import (
	"team-management/members/domain"
	"team-management/members/utils"
)

const idField = "_id"

type Database interface {
	Get(fieldName, fieldVale string, output interface{}) error
	Create(content interface{}) (string, error)
	Update(fieldName, fieldVale string, content interface{}) (bool, error)
	Delete(fieldName, fieldVale string) (bool, error)
	Filter(filters map[string]string, output interface{}) error
}

type MembersRepository struct {
	db Database
}

func NewMembersRepository(db Database) *MembersRepository {
	return &MembersRepository{
		db: db,
	}
}

func (m *MembersRepository) Get(id string) (*domain.Member, error) {
	var memberDB MemberDB
	if err := m.db.Get(idField, id, &memberDB); err != nil {
		return &domain.Member{}, err
	}
	return NewMemberDomain(&memberDB), nil
}

func (m *MembersRepository) Create(member *domain.Member) (string, error) {
	memberDB := NewMemberDB(member)
	return m.db.Create(memberDB)
}

func (m *MembersRepository) Delete(id string) (bool, error) {
	return m.db.Delete(idField, id)
}

func (m *MembersRepository) Update(member *domain.Member) (bool, error) {
	memberDB := NewMemberDB(member)
	return m.db.Update(idField, member.ID, memberDB)
}

func (m *MembersRepository) Filter(filters map[string]string) ([]*domain.Member, error) {
	validFilters := map[string]bool{"name": true, "type": true, "tags": true}

	members := []*domain.Member{}
	membersDB := []*MemberDB{}

	for field := range filters {
		if _, exists := validFilters[field]; !exists {
			return members, utils.NewBadRequest("The following filter field is invalid: " + field)
		}
	}

	if err := m.db.Filter(filters, &membersDB); err != nil {
		return make([]*domain.Member, 0), err
	}

	for _, memberDB := range membersDB {
		members = append(members, NewMemberDomain(memberDB))
	}

	return members, nil
}
