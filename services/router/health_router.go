package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ferdn4ndo/userver-logger-api/controllers"
)

func HealthRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", controllers.GetHealthState)

	return router
}
