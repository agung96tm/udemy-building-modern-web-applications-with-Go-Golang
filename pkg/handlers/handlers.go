package handlers

import (
	"net/http"

	"github.com/agung96tm/udemy-modern-go/pkg/config"
	"github.com/agung96tm/udemy-modern-go/pkg/models"
	"github.com/agung96tm/udemy-modern-go/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// HomeHandler root of page
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	_ = render.RenderTemplate(w, "home.page", &models.TemplateData{})
}

// AboutHandler about page handler
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello About!!!"

	_ = render.RenderTemplate(w, "about.page", &models.TemplateData{
		StringMap: stringMap,
	})
}
