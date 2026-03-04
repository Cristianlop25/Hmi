package main

import (
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, name string, data any) {
	tpl := template.Must(template.ParseFiles(
		"views/layouts/base.html",
		"views/pages/"+name+".html",
	))

	if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

type App struct {
	render func(http.ResponseWriter, string, any)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		app.render(w, "connectors", nil)
	case "/connectors":
		app.render(w, "connectors", nil)
	case "/identification":
		app.render(w, "identification", nil)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app.css", http.FileServer(http.Dir("public")))

	app := &App{render: render}
	mux.Handle("/", app)

	log.Println("http://localhost:8081")
	_ = http.ListenAndServe("localhost:8081", mux)
}
