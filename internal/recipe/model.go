package recipe

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/heyaustria/recipe-manager/internal/category"
)

type Model struct {
	ID         uint64            `json:"id"`
	Name       string            `json:"name"`
	Rating     int               `json:"rating"`
	Categories []*category.Model `json:"categories"`
}

func (m *Model) Bind(r *http.Request) error {
	m.ID = uint64(rand.Intn(100) + 10)
	return nil
}

func (m *Model) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
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

type List struct {
	Total   int      `json:"total"`
	Recipes []*Model `json:"items"`
}

func (m *List) Bind(r *http.Request) error {
	return nil
}

func (m *List) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
