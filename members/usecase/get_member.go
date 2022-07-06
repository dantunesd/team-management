package usecase

import (
	"encoding/json"
	"team-management/members/domain"
	"team-management/members/utils"
)

type GetMember struct {
	repository domain.MembersRepository
}

func NewGetMember(repository domain.MembersRepository) *GetMember {
	return &GetMember{
		repository: repository,
	}
}

type GetMemberInput struct {
	ID string
}

type GetMemberOutput struct {
	ID        string
	Name      string
	Type      string
	TypeData  json.RawMessage
	Tags      []string
	CreatedAt string
	UpdatedAt string
}

func (g *GetMember) Handle(id string) (*GetMemberOutput, error) {
	member, err := g.repository.Get(id)
	if err != nil {
		return &GetMemberOutput{}, err
	}

	if member.ID == "" {
		return &GetMemberOutput{}, utils.NewNotFound("member not found")
	}

	convertedTypeData, _ := json.Marshal(member.TypeData)

	return &GetMemberOutput{
		ID:        member.ID,
		Name:      member.Name,
		Type:      string(member.Type),
		TypeData:  convertedTypeData,
		Tags:      member.Tags,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
	}, nil
}
