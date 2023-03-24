package category

import (
	"errors"
	"math/rand"
	"net/http"
)

type Model struct {
	ID   int64
	Name string
}

func (m *Model) ValidateName() error {
	if m.Name == "" {
		return errors.New("name cannot be empty")
	}
	if len(m.Name) > 255 {
		return errors.New("name is too long, can only be 255 characters or less")
	}
	return nil
}

func (m *Model) Bind(r *http.Request) error {
	m.ID = int64(rand.Intn(100) + 10)
	return nil
}

func (m *Model) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type List struct {
	Total int      `json:"total"`
	Items []*Model `json:"items"`
}

func (m *List) Bind(r *http.Request) error {
	return nil
}

func (m *List) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
