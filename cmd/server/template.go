package server

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func newTemplate() *Template {
	templates, err := template.ParseGlob("./public/template/*.html") // 修改路径为实际模板路径
	if err != nil {
		log.Fatal("Failed to parse template:", err)
	}
	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
