package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/s3niffer/taskmanagementapp/internal/app"
)

func SetUpHandler(a *app.Application) *chi.Mux {
	handler := chi.NewRouter()

	handler.Get("/health", a.HealthCheck)

	handler.Get("/users/{id}", a.User.HandleGetUserByID)
	handler.Post("/users", a.User.HandleCreateUser)

	return handler
}
