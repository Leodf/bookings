package main

import (
	"net/http"

	"github.com.br/Leodf/bookings/internal/config"
	"github.com.br/Leodf/bookings/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(_ *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/generals-quarters", handler.Repo.Generals)
	mux.Get("/majors-suite", handler.Repo.Majors)

	mux.Get("/search-availability", handler.Repo.Availability)
	mux.Post("/search-availability", handler.Repo.PostAvailability)
	mux.Post("/search-availability-json", handler.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handler.Repo.ChooseRoom)

	mux.Get("/contact", handler.Repo.Contact)

	mux.Get("/make-reservation", handler.Repo.Reservation)
	mux.Post("/make-reservation", handler.Repo.PostReservation)
	mux.Get("/reservation-summary", handler.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
