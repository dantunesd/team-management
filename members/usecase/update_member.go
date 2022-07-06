package usecase

import (
	"encoding/json"
	"team-management/members/domain"
	"team-management/members/utils"
)

type UpdateMember struct {
	repository domain.MembersRepository
}

func NewUpdateMember(repository domain.MembersRepository) *UpdateMember {
	return &UpdateMember{
		repository: repository,
	}
}

type UpdateMemberInput struct {
	Name     string
	Type     string
	TypeData json.RawMessage
	Tags     []string
}

type UpdateMemberOutput struct {
	UpdatedAt string
}

func (u *UpdateMember) Handle(id string, input *UpdateMemberInput) (*UpdateMemberOutput, error) {
	updateMemberOutput := &UpdateMemberOutput{}

	typeData, err := domain.TypeDataFactory(domain.Type(input.Type), input.TypeData)
	if err != nil {
		return updateMemberOutput, err
	}

	member := &domain.Member{
		ID:        id,
		Name:      input.Name,
		Type:      domain.Type(input.Type),
		TypeData:  typeData,
		Tags:      input.Tags,
		UpdatedAt: utils.GetTimeNowInUTC(),
	}

	if err := member.Validate(); err != nil {
		return updateMemberOutput, err
	}

	updated, err := u.repository.Update(member)
	if err != nil {
		return updateMemberOutput, err
	}

	if !updated {
		return updateMemberOutput, utils.NewNotFound("member not found")
	}

	updateMemberOutput.UpdatedAt = member.UpdatedAt

	return updateMemberOutput, nil
}
