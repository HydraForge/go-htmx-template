package main

import (
	"fmt"
	"hydraforge/go-htmx-template/frontend/src/components"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(components.Home()).ServeHTTP)

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("works")
}
