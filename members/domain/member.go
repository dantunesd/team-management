package domain

import (
	"team-management/members/utils"
)

type Tags []string

type Member struct {
	ID        string
	Name      string
	Type      Type
	TypeData  TypeData
	Tags      Tags
	CreatedAt string
	UpdatedAt string
}

func (m *Member) Validate() error {
	if !m.isValidType() {
		return utils.NewBadRequest("Type and type_data should match")
	}

	if !m.isValidTypeData() {
		return utils.NewBadRequest("The field type_data has invalid fields due to type")
	}

	return nil
}

func (m *Member) isValidType() bool {
	return m.Type == m.TypeData.GetType()
}

func (m *Member) isValidTypeData() bool {
	return m.TypeData.IsValid()
}
