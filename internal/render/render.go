package render

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, name string, data any) {
	if r.Header.Get("HX-Request") == "true" {
		tpl := template.Must(template.ParseFiles(
			"views/pages/" + name + ".html",
		))

		_ = tpl.ExecuteTemplate(w, "content", data)
		return
	}

	tpl := template.Must(template.ParseFiles(
		"views/layouts/base.html",
		"views/pages/"+name+".html",
	))

	_ = tpl.ExecuteTemplate(w, "base", data)
}
