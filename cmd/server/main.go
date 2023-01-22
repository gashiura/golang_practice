package main

import (
	"net/http"
	"time"

	"golang_practice/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Start server.")
	r := chi.NewRouter()
	r.Use(middleware.Timeout(1 * time.Second))
	r.Route("/users", func(r chi.Router) {
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", controller.UserShowAction)
		})
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("Can't start server.")
	}
}
