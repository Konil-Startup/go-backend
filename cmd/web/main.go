package main

import (
	"log"
	"net/http"

	"github.com/Konil-Startup/go-backend/pkg/config"
	"github.com/Konil-Startup/go-backend/pkg/handlers"
	"github.com/Konil-Startup/go-backend/pkg/render"
)

var (
	port = ":8080"
)

func main() {
	var app config.AppConfig

	tCache, err := render.CacheTemplates()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tCache
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	log.Printf("Starting server on port %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}
