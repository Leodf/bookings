package main

import (
	"net/http"

	"github.com.br/Leodf/bookings/pkg/config"
	"github.com.br/Leodf/bookings/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	mux.Get("/generals-quarters", http.HandlerFunc(handler.Repo.Generals))
	mux.Get("/majors-suite", http.HandlerFunc(handler.Repo.Majors))
	mux.Get("/search-availability", http.HandlerFunc(handler.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handler.Repo.PostAvailability))
	mux.Get("/contact", http.HandlerFunc(handler.Repo.Contact))

	mux.Get("/make-reservation", http.HandlerFunc(handler.Repo.Reservation))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
