package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/s3niffer/taskmanagementapp/internal/app"
)

func SetUpHandler(a *app.Application) *chi.Mux {
	handler := chi.NewRouter()

	handler.Get("/health", a.HealthCheck)

	handler.Route("/users", func(r chi.Router) {
		r.Get("/{id}", a.User.HandleGetUserByID)
		r.Post("/", a.User.HandleCreateUser)
	})

	return handler
}
