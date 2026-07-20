package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/s3niffer/taskmanagementapp/internal/app"
)

func New(app app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	r.Post("/register", app.UserApi.RegisterUser)

	return r
}
