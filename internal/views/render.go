package render

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, name string, data any) {
	tpl := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/"+name+".html",
	))

	tpl = template.Must(tpl.ParseGlob("templates/components/*.html"))

	if r.Header.Get("HX-Request") == "true" {
		_ = tpl.ExecuteTemplate(w, "content", data)
		return
	}

	_ = tpl.ExecuteTemplate(w, "base", data)
}
