package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ferdn4ndo/userver-logger-api/controllers"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/render"
)

func HealthRouter(dbService database.DatabaseServiceInterface, renderService render.RenderServiceInterface) chi.Router {
	router := chi.NewRouter()

	controller := controllers.HealthController{DbService: dbService, RenderService: renderService}

	router.Get("/", controller.GetHealthState)

	return router
}
