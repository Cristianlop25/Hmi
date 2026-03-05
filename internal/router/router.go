package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID    int
	Title string
}

type Router struct {
	render func(http.ResponseWriter, *http.Request, string, any)
}

func New(render func(http.ResponseWriter, *http.Request, string, any)) http.Handler {
	r := &Router{render: render}
	router := chi.NewRouter()

	router.Get("/", r.connectors)
	router.Get("/connectors", r.connectors)
	router.Get("/identification", r.identification)

	return router
}

func (r *Router) connectors(w http.ResponseWriter, req *http.Request) {
	data := map[string]any{
		"Title": "Connectors",
	}
	r.render(w, req, "connectors", data)
}

func (r *Router) identification(w http.ResponseWriter, req *http.Request) {
	posts := []Post{
		{
			ID:    1,
			Title: "Identification 1",
		},
		{
			ID:    2,
			Title: "Identification 2",
		},
	}
	data := map[string]any{
		"Title": "Identification",
		"Year":  time.Now().Year(),
		"Posts": posts,
	}
	r.render(w, req, "identification", data)
}
