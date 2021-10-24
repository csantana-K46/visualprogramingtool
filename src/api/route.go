package main

import "github.com/go-chi/chi"
import c "api/controllers"

func GetRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", c.Get)
	return r
}