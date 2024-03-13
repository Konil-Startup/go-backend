package handlers

import (
	"fmt"
	"net/http"

	"github.com/Konil-Startup/go-backend/pkg/config"
	"github.com/Konil-Startup/go-backend/pkg/models"
	"github.com/Konil-Startup/go-backend/pkg/render"
	"github.com/brianvoe/gofakeit"
)

var Repo *Repository

var stringMap = []map[string]interface{}{
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/9.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
		"HelpCount":  10,
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/1.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
		"HelpCount":  100,
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/34.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/19.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
		"HelpCount":  99,
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/20.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/21.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/22.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/23.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
		"HelpCount":  10,
	},
	{
		"ImageLink":  "https://randomuser.me/api/portraits/men/24.jpg",
		"Title":      gofakeit.Name(),
		"Company":    gofakeit.Company(),
		"Occupation": fmt.Sprintf("%s %s %s", gofakeit.Job().Level, gofakeit.Job().Descriptor, gofakeit.Job().Title),
		"Years":      int(gofakeit.Price(1, 10)),
		"Price":      int(gofakeit.Price(1000, 10000)),
	},
}

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
	}

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{Data: stringMap})
}

func (m *Repository) Show(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{Data: []map[string]interface{}{
		{
			"ImageLink": r.FormValue("image"),
			"Title":     r.FormValue("title"),
			"Com":       r.FormValue("com"),
			"Job":       r.FormValue("job"),
			"Years":     r.FormValue("years"),
			"Price":     r.FormValue("price"),
		},
	}})
}
