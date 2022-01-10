package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/igetum/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"

	"github.com/igetum/bookings/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Use(NoServe)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
