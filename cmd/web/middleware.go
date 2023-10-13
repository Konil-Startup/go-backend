package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit page at", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteDefaultMode,
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit page at", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
