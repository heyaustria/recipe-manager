package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/heyaustria/recipe-manager/internal/category"
	"github.com/heyaustria/recipe-manager/internal/recipe"
	"github.com/heyaustria/recipe-manager/internal/rest"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Route("/recipes", func(r chi.Router) {
		handler := rest.NewRecipeHandler(recipe.NewService(recipe.NewRepository(), category.NewService(category.NewRepository())))
		r.Post("/", handler.Create)
		r.Get("/", handler.List)
		r.Get("/{id}", handler.Read)
		r.Put("/{id}", handler.Update)
	})

	r.Route("/categories", func(r chi.Router) {
		handler := rest.NewCategoryHandler(category.NewService(category.NewRepository()))
		r.Post("/", handler.Create)
		r.Get("/", handler.List)
		r.Get("/{id}", handler.Read)
		r.Put("/{id}", handler.Update)
	})

	http.ListenAndServe(":3000", r)
}
