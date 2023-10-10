package main

import (
	"log"
	"net/http"
)

var (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting server on port %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}
