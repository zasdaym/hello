package http

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func respondError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	if code == http.StatusInternalServerError {
		log.Error().Err(err).Send()
	}
}

func respondJSON(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
