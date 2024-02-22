package main

import (
	"fmt"
	"hydraforge/go-htmx-template/frontend/src/views/home"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(home.Home()).ServeHTTP)
	r.Get("/test", templ.Handler(home.Test()).ServeHTTP)

	fileServer := http.FileServer(http.Dir("./frontend/src/assets/"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	fmt.Println("running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
