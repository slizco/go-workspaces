package main

import (
	"log"
	"net/http"
	"os"
	"sample/blue"
)

func main() {
	http.Handle("/", &blue.Handler{})

	port := os.Getenv("PORT")

	if port == "" {
		log.Print("$PORT is unset, using 8080")
		port = "8080"
	}

	log.Printf("Starting server, listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
