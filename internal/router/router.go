package router

import "net/http"

type Router struct {
}

func CreateRouter() *Router {
	return &Router{}
}

func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		{
			w.Write([]byte("hello"))
		}
	case "/2":
		{
			w.Write([]byte("bye"))
		}
	default:
		w.Write([]byte("help"))
	}
}
