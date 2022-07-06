package domain

type MembersRepositoryReader interface {
	Get(id string) (*Member, error)
}

type MembersRepositoryWritter interface {
	Create(member *Member) (string, error)
	Update(member *Member) (bool, error)
	Delete(id string) (bool, error)
	Filter(filters map[string]string) ([]*Member, error)
}

type MembersRepository interface {
	MembersRepositoryReader
	MembersRepositoryWritter
}
