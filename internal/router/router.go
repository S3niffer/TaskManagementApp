package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/s3niffer/taskmanagementapp/internal/app"
)

func New(app app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/users", app.Handler.CreateUser)
	r.Get("/users", app.Handler.GetAllUsers)

	r.Get("/health", app.HealthCheck)
	return r
}
