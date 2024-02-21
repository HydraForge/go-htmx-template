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

	fileServer := http.FileServer(http.Dir("./frontend/src/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	fmt.Println("running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
