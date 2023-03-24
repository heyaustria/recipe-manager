package recipe

import "github.com/heyaustria/recipe-manager/internal/category"

type service struct {
	repo     Repository
	category category.Service
}

func NewService(r Repository, c category.Service) *service {
	return &service{
		repo:     r,
		category: c,
	}
}

func (s *service) Create(m *Model) (*Model, error) {
	if err := m.ValidateName(); err != nil {
		return nil, err
	}
	if m.Categories != nil {
		s.category.Exists(m.Categories)
	}
	return s.repo.Create(m)
}

func (s *service) Read(id uint64) (*Model, error) {
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
