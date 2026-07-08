package router

import (
	"net/http"

	"github.com/s3niffer/taskmanagementapp/internal/app"
)

type Router struct {
	App *app.Application
}

func New(app *app.Application) *Router {
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
	default:
		w.Write([]byte("help"))
	}
}
