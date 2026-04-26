package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/agung96tm/udemy-go-bookings/pkg/config"
	"github.com/agung96tm/udemy-go-bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplate(cfg *config.AppConfig) {
	app = cfg
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpt string, data *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpt]
	if !ok {
		return fmt.Errorf("template %q not found", tmpt)
	}

	data = AddDefaultData(data)
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("templates/*.page.gohtml")
	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		baseName := filepath.Base(page)
		// Handlers use short names like "home" for home.page.gohtml
		name := strings.TrimSuffix(baseName, ".gohtml")
		tmpl, err := template.New(baseName).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.gohtml")
		if err != nil {
			return tmplCache, err
		}
		if len(matches) > 0 {
			tmpl, err = tmpl.ParseGlob("templates/*.layout.gohtml")
			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[name] = tmpl
	}

	return tmplCache, nil
}
