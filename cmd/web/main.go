package main

import (
	"log"
	"net/http"

	"github.com/Konil-Startup/go-backend/pkg/handlers"
)

var (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting server on port %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}
