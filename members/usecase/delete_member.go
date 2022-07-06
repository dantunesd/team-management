package usecase

import (
	"team-management/members/domain"
	"team-management/members/utils"
)

type DeleteMember struct {
	repository domain.MembersRepository
}

func NewDeleteMember(repository domain.MembersRepository) *DeleteMember {
	return &DeleteMember{
		repository: repository,
	}
}

type DeleteMemberInput struct {
	ID string
}

func (g *DeleteMember) Handle(id string) error {
	wasDeleted, err := g.repository.Delete(id)
	if err != nil {
		return err
	}

	if !wasDeleted {
		return utils.NewNotFound("member not found")
	}

	return nil
}
