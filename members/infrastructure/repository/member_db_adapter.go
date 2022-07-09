package repository

import (
	"encoding/json"
	"team-management/members/domain"
)

type MemberDB struct {
	ID        string   `bson:"_id,omitempty"`
	Name      string   `bson:"name,omitempty"`
	Type      string   `bson:"type,omitempty"`
	TypeData  []byte   `bson:"type_data,omitempty"`
	Tags      []string `bson:"tags,omitempty"`
	CreatedAt string   `bson:"created_at,omitempty"`
	UpdatedAt string   `bson:"updated_at,omitempty"`
}

func NewMemberDB(member *domain.Member) *MemberDB {
	typeData, _ := json.Marshal(member.TypeData)

	return &MemberDB{
		Name:      member.Name,
		Type:      string(member.Type),
		TypeData:  typeData,
		Tags:      member.Tags,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	}
}

func NewMemberDomain(memberDB *MemberDB) *domain.Member {
	typeData, _ := domain.TypeDataFactory(domain.Type(memberDB.Type), memberDB.TypeData)

	return &domain.Member{
		ID:        memberDB.ID,
		Name:      memberDB.Name,
		Type:      domain.Type(memberDB.Type),
		TypeData:  typeData,
		Tags:      memberDB.Tags,
		CreatedAt: memberDB.CreatedAt,
		UpdatedAt: memberDB.UpdatedAt,
	}
}
