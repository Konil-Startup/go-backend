package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port = ":8080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi from server", time.Now())
	})

	log.Printf("Starting server on port %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}
