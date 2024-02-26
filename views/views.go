package views

import (
	"html/template"
	"path"
)

func RenderPage(file string) *template.Template {
	layout := path.Join("views", "layouts", "baseLayout.html")
	navbar := path.Join("views", "components", "navbar.html")
	footer := path.Join("views", "components", "footer.html")
	return template.Must(template.ParseFiles(layout, navbar, footer, file))
}

func Render404() *template.Template {
	page := path.Join("views", "404.html")
	tmpl := RenderPage(page)
	return tmpl
}
