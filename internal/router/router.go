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
	render func(http.ResponseWriter, string, any)
}

func New(render func(http.ResponseWriter, string, any)) http.Handler {
	r := &Router{render: render}
	router := chi.NewRouter()

	router.Get("/", r.connectors)
	router.Get("/connectors", r.connectors)
	router.Get("/identification", r.identification)

	return router
}

func (r *Router) connectors(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"Title": "Anasayfa",
		"User":  "Uygar",
	}
	r.render(w, "connectors", data)
}

func (r *Router) identification(w http.ResponseWriter, _ *http.Request) {
	posts := []Post{
		{
			ID:    1,
			Title: "Go Templates ile Başlarken",
		},
		{
			ID:    2,
			Title: "HTMX’a Giriş",
		},
	}

	data := map[string]any{
		"Title": "Identification",
		"Year":  time.Now().Year(),
		"Posts": posts,
	}
	r.render(w, "identification", data)
}
