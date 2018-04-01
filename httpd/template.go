package httpd

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// NewTemplate returns an instance of Template responsible for html templates
// at path.
func NewTemplate(path string) *Template {
	t := template.Must(template.ParseGlob(path))
	fmt.Printf("%+v\n", t.DefinedTemplates())
	return &Template{
		templates: t,
	}
}

// Template manages html templates.
type Template struct {
	templates *template.Template
}

// Render implements the echo.Renderer interface.
func (t *Template) Render(w io.Writer, name string, data interface{},
	c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
