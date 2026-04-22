package main

import (
	"net/http"

	"github.com/agung96tm/udemy-modern-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *Application) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(NoSurf)

	r.Get("/", handlers.Repo.HomeHandler)
	r.Get("/about", handlers.Repo.AboutHandler)

	return r
}
