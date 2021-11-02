package route

import (
	controller "api/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func GetRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Get("/", controller.Get)
	r.Post("/api/program/parse", controller.ParseCode)
	r.Post("/api/program/run", controller.RunCode)
	r.Get("/api/programs", controller.List)
	r.Post("/api/program/add", controller.Add)
	r.Get("/api/programs/{id}", controller.GetDetail)
	return r
}
