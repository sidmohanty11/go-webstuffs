package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/config"
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/rooms/generals", handlers.Repo.Generals)
	mux.Get("/rooms/presidentals", handlers.Repo.Presidentals)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/signup", handlers.Repo.SignUp)
	mux.Post("/signup", handlers.Repo.PostSignUp)
	mux.Get("/login", handlers.Repo.Login)
	mux.Get("/logout", handlers.Repo.Logout)
	mux.Post("/login", handlers.Repo.PostLogin)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handlers.Repo.AdminDashboard)
		mux.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		mux.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		mux.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
	})

	fileServer := http.FileServer(http.Dir("./static/")) //for serving static files
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
