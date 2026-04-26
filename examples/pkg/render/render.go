package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/agung96tm/udemy-modern-go/pkg/config"
	"github.com/agung96tm/udemy-modern-go/pkg/models"
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

//
//func RenderTemplatett(w http.ResponseWriter, tmpl string, data interface{}) error {
//	t, err := template.ParseFiles(fmt.Sprintf("./templates/%s.gohtml", tmpl), "./templates/base.layout.gohtml")
//	if err != nil {
//		return err
//	}
//	return t.Execute(w, data)
//}
//
//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string, data interface{}) error {
//	var tmpl *template.Template
//	var err error
//
//	_, inMap := tc[t]
//	if !inMap {
//		log.Printf("creating template and cached: %v\n", t)
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Printf("Error creating template cache: %v\n", err)
//		}
//	} else {
//		log.Println("using cached template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, data)
//	if err != nil {
//		log.Printf("Error executing template: %v\n", err)
//		return err
//	}
//
//	return nil
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s.gohtml", t),
//		"./templates/base.layout.gohtml",
//	}
//
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//
//	tc[t] = tmpl
//
//	return nil
//}
