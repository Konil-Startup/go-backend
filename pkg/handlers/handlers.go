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
	render.RenderTemplate(w, "home.page.tmpl", nil)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := map[string]interface{}{}
	stringMap["test"] = "Hello"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{Data: stringMap})
}
