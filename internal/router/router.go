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

type RenderFunc func(http.ResponseWriter, *http.Request, string, any)

type Router struct {
	render RenderFunc
}

type ConnectorType string

const (
	CCS2    ConnectorType = "CCS2"
	CCS1    ConnectorType = "CCS1"
	Schuko  ConnectorType = "Schuko"
	CHAdeMO ConnectorType = "CHAdeMO"
	Type1   ConnectorType = "Type1"
	Type2   ConnectorType = "Type2"
)

type Connector struct {
	Name string
	Type ConnectorType
}

func (connectorType ConnectorType) IconPath() string {
	connectorsDir := "/assets/images/connectors/"
	switch connectorType {
	case CCS2:
		return connectorsDir + "ccs2.svg"
	case CCS1:
		return connectorsDir + "ccs1.svg"
	case Schuko:
		return connectorsDir + "schuko.svg"
	case Type1:
		return connectorsDir + "type1.svg"
	case Type2:
		return connectorsDir + "type2.svg"
	case CHAdeMO:
		return connectorsDir + "chademo.svg"
	default:
		return connectorsDir + "default.svg"
	}
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
	connectors := []Connector{
		{Name: "Socket A", Type: CCS1},
		{Name: "Socket B", Type: CCS2},
	}
	data := map[string]any{
		"Connectors": connectors,
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
