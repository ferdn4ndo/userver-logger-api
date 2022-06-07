package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/logging"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewHandler() *chi.Mux {
	router := chi.NewRouter()
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	return router
}

func UnauthorizedHandler(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(ErrUnauthorized.StatusCode)

	err := json.NewEncoder(writer).Encode(ErrUnauthorized)
	if err != nil {
		logging.Errorf("Error encoding json object: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "An internal error occurred while encoding the json object!")
		return
	}
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(405)
	render.Render(writer, request, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(404)
	render.Render(w, r, ErrNotFound)
}

type ControllerHandler struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
