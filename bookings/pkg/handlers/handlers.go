package handlers

import (
	"net/http"

	"github.com/agung96tm/udemy-go-bookings/pkg/config"
	"github.com/agung96tm/udemy-go-bookings/pkg/models"
	"github.com/agung96tm/udemy-go-bookings/pkg/render"
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	m.App.Session.Put(r.Context(), "name", "agung")
	_ = render.RenderTemplate(w, "home.page", &models.TemplateData{})
}

// AboutHandler about page handler
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello About!!!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	name := m.App.Session.GetString(r.Context(), "name")
	stringMap["name"] = name

	_ = render.RenderTemplate(w, "about.page", &models.TemplateData{
		StringMap: stringMap,
	})
}
