package router

import (
	"net/http"

	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/handler"
)

type Router struct {
	App app.Application
}

func New(app app.Application) *Router {
	return &Router{
		App: app,
	}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		w.Write([]byte("hello"))
	case "/2":
		w.Write([]byte("bye"))
	case "/health":
		router.App.HealthCheck(w, r)
	case "/user":
		handler.CreateUser(router.App, w, r)
	default:
		w.Write([]byte("help"))
	}
}
