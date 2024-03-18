package helpers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type templ struct {
	templates *template.Template
}

var Templs = &templ{
	templates: findAndParseTemplates("./web/views/", nil),
}

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) *template.Template {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, "html") {
			if e1 != nil {
				return e1
			}

			b, e2 := os.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := strings.ReplaceAll(path[pfx:], "\\", "/")
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})
	fmt.Printf("defined: %v\n", root.DefinedTemplates())
	if err != nil {
		panic(err)
	}
	return root
}
func (t *templ) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
