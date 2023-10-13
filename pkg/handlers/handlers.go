package handlers

import (
	"net/http"

	"github.com/Konil-Startup/go-backend/pkg/config"
	"github.com/Konil-Startup/go-backend/pkg/models"
	"github.com/Konil-Startup/go-backend/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	if !m.App.Session.Exists(r.Context(), "remote_ip") {
		m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	} else {
		m.App.Session.Remove(r.Context(), "remote_ip")
	}
	render.RenderTemplate(w, "home.page.tmpl", nil)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")
	stringMap := map[string]interface{}{}
	stringMap["test"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{Data: stringMap})
}
