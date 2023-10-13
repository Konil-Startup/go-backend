package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Konil-Startup/go-backend/pkg/config"
	"github.com/Konil-Startup/go-backend/pkg/handlers"
	"github.com/Konil-Startup/go-backend/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const (
	port = ":8080"
)

var (
	session *scs.SessionManager
	app     config.AppConfig
)

func main() {

	app.InProduction = false

	tCache, err := render.CacheTemplates()
	if err != nil {
		log.Fatal(err)
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = true
	app.Session = session

	app.TemplateCache = tCache
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Printf("Starting server on port %v\n", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
