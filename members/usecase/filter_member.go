package usecase

import (
	"encoding/json"
	"team-management/members/domain"
)

type FilterMember struct {
	repository domain.MembersRepository
}

func NewFilterMember(repository domain.MembersRepository) *FilterMember {
	return &FilterMember{
		repository: repository,
	}
}

type FilterMemberInput struct {
	Filters map[string]string
}

type FilterMemberOutput struct {
	ID        string
	Name      string
	Type      string
	TypeData  json.RawMessage
	Tags      []string
	CreatedAt string
	UpdatedAt string
}

func (f *FilterMember) Handle(input *FilterMemberInput) ([]FilterMemberOutput, error) {
	filterMemberOutput := []FilterMemberOutput{}

	members, err := f.repository.Filter(input.Filters)
	if err != nil {
		return filterMemberOutput, err
	}

	for _, member := range members {
		convertedTypeData, _ := json.Marshal(member.TypeData)
		output := FilterMemberOutput{
			ID:        member.ID,
			Name:      member.Name,
			Type:      string(member.Type),
			TypeData:  convertedTypeData,
			Tags:      member.Tags,
			CreatedAt: member.CreatedAt,
			UpdatedAt: member.UpdatedAt,
		}
		filterMemberOutput = append(filterMemberOutput, output)
	}

	return filterMemberOutput, nil
}
