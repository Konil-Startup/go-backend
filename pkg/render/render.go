package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Konil-Startup/go-backend/pkg/config"
)

var (
	functions = template.FuncMap{}
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CacheTemplates()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not find template:", tmpl)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CacheTemplates() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
