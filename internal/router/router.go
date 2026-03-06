package router

import (
	"net/http"

	"hmi-sonic/internal/connectors"
	"hmi-sonic/internal/identification"

	"github.com/go-chi/chi/v5"
)

type RenderFunc func(http.ResponseWriter, *http.Request, string, any)

type Router struct {
	render                RenderFunc
	connectorsService     connectors.Service
	identificationService identification.Service
}

func New(render RenderFunc, connectorsService connectors.Service, identificationService identification.Service) http.Handler {
	r := &Router{
		render:                render,
		connectorsService:     connectorsService,
		identificationService: identificationService,
	}
	router := chi.NewRouter()

	router.Get("/", r.connectors)
	router.Get("/connectors", r.connectors)
	router.Get("/identification", r.identification)

	return router
}

func (router *Router) connectors(writer http.ResponseWriter, request *http.Request) {
	connectors, err := router.connectorsService.List()
	if err != nil {
		http.Error(writer, "failed to load connectors", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"Connectors": connectors,
	}
	router.render(writer, request, "connectors", data)
}

func (router *Router) identification(writer http.ResponseWriter, request *http.Request) {
	identification, err := router.identificationService.Status()
	if err != nil {
		http.Error(writer, "failed to load identification", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"Identification": identification,
	}
	router.render(writer, request, "identification", data)
}
