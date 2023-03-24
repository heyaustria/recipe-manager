package category

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) Create(m *Model) (*Model, error) {
	if err := m.ValidateName(); err != nil {
		return nil, err
	}
	return s.repo.Create(m)
}

func (s *service) Read(id int64) (*Model, error) {
	return s.repo.Read(id)
}

func (s *service) Update(m *Model) (*Model, error) {
	if err := m.ValidateName(); err != nil {
		return nil, err
	}
	return s.repo.Update(m)
}

func (s *service) List() (*List, error) {
	return s.repo.List()
}

func (s *service) Exists(cats []*Model) error {
	return s.repo.Exists(cats)
}
