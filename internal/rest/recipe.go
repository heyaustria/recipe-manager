package rest

import (
	"net/http"
	"strconv"

	"github.com/heyaustria/recipe-manager/internal/recipe"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type RecipeHandler struct {
	service recipe.Service
}

func NewRecipeHandler(s recipe.Service) *RecipeHandler {
	return &RecipeHandler{
		service: s,
	}
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := &recipe.Model{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	recipe, err := h.service.Create(data)
	if err != nil {
		// handle error
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, recipe)
}

func (h *RecipeHandler) Read(w http.ResponseWriter, r *http.Request) {
	var recipe *recipe.Model
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	recipe, err = h.service.Read(uint64(intID))
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, recipe)
}

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	data := &recipe.Model{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	data.ID = uint64(intID)
	recipe, err := h.service.Update(data)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, recipe)
}

func (h *RecipeHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List()
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, list)
}
