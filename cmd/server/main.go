package main

import (
	"net/http"

	"golang_practice/controller"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Start server.")

	if err := http.ListenAndServe(":8080", controller.Routes()); err != nil {
		log.Fatal().Err(err).Msg("Can't start server.")
	}
}
