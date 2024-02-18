package http

import (
	"net/http"
	"strings"
)

func echoHandler() http.HandlerFunc {
	type response struct {
		Headers map[string]string `json:"headers"`
		Path    string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		headers := make(map[string]string)
		for k, v := range r.Header {
			headers[k] = strings.Join(v, ",")
		}
		respondJSON(w, response{
			Headers: headers,
			Path:    r.URL.Path,
		})
	}
}
