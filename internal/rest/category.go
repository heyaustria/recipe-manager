package rest

import (
	"net/http"
	"strconv"

	"github.com/heyaustria/recipe-manager/internal/category"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CategoryHandler struct {
	service category.Service
}

func NewCategoryHandler(s category.Service) *CategoryHandler {
	return &CategoryHandler{
		service: s,
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	data := &category.Model{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	cat, err := h.service.Create(data)
	if err != nil {
		// handle error
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, cat)
}

func (h *CategoryHandler) Read(w http.ResponseWriter, r *http.Request) {
	var cat *category.Model
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	cat, err = h.service.Read(int64(intID))
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, cat)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	data := &category.Model{}
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
	data.ID = int64(intID)
	cat, err := h.service.Update(data)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, cat)
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List()
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, list)
}
