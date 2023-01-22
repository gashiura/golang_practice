package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang_practice/model"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func UserShowAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		log.Error().Err(err).Msg("userID is not integer.")
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	user := model.User{ID: id, UID: "TestUser1", Email: "example+1000@gmail.com"}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Error().Err(err).Msg("user not found.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
