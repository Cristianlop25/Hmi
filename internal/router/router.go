package router

import (
	"net/http"
	"time"

	"hmi-sonic/internal/connectors"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID    int
	Title string
}

type RenderFunc func(http.ResponseWriter, *http.Request, string, any)

type Router struct {
	render            RenderFunc
	connectorsService connectors.Service
}

func New(render RenderFunc, connectorsService connectors.Service) http.Handler {
	r := &Router{
		render:            render,
		connectorsService: connectorsService,
	}
	router := chi.NewRouter()

	router.Get("/", r.connectors)
	router.Get("/connectors", r.connectors)
	router.Get("/identification", r.identification)

	return router
}

func (router *Router) connectors(w http.ResponseWriter, req *http.Request) {
	connectors, err := router.connectorsService.List()
	if err != nil {
		http.Error(w, "failed to load connectors", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"Connectors": connectors,
	}
	router.render(w, req, "connectors", data)
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
