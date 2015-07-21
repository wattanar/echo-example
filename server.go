package main

import (
	"io"
	"net/http"
	"html/template"
	"github.com/labstack/echo"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c *echo.Context) error {
    return c.Render(http.StatusOK, "hello", "World")
}

func main() {

	e := echo.New()

	t := &Template{
	    templates: template.Must(template.ParseGlob("public/*.html")),
	}

	e.SetRenderer(t)

	e.Get("/", Hello)

	e.Run(":1323")
	
}