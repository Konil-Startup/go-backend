package main

import (
	"net/http"

	"github.com/Konil-Startup/go-backend/pkg/config"
	"github.com/Konil-Startup/go-backend/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer, NoSurf, middleware.Logger, LoadSession)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
