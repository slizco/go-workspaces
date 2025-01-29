package main

import (
	"log"
	"net/http"
	"sample/blue"
)

func main() {
	http.Handle("/", &blue.Handler{})

	log.Print("Starting server, listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
