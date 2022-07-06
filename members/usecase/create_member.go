package usecase

import (
	"encoding/json"
	"team-management/members/domain"
	"team-management/members/utils"
)

type CreateMember struct {
	repository domain.MembersRepository
}

func NewCreateMember(repository domain.MembersRepository) *CreateMember {
	return &CreateMember{
		repository: repository,
	}
}

type CreateMemberInput struct {
	Name     string
	Type     string
	TypeData json.RawMessage
	Tags     []string
}

type CreateMemberOutput struct {
	ID        string
	CreatedAt string
}

func (c *CreateMember) Handle(input *CreateMemberInput) (*CreateMemberOutput, error) {
	typeData, err := domain.TypeDataFactory(domain.Type(input.Type), input.TypeData)
	if err != nil {
		return &CreateMemberOutput{}, err
	}

	member := &domain.Member{
		Name:      input.Name,
		Type:      domain.Type(input.Type),
		TypeData:  typeData,
		Tags:      input.Tags,
		CreatedAt: utils.GetTimeNowInUTC(),
	}

	if err := member.Validate(); err != nil {
		return &CreateMemberOutput{}, err
	}

	id, err := c.repository.Create(member)
	if err != nil {
		return &CreateMemberOutput{}, err
	}

	member.ID = id

	return &CreateMemberOutput{
		ID:        member.ID,
		CreatedAt: member.CreatedAt,
	}, nil
}
