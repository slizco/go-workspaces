package blue

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1 style="color:DodgerBlue;">Hello World</h1>`)
}
