package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	"github.com/ferdn4ndo/userver-logger-api/controllers"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
)

func LogEntryRouter(dbService database.DatabaseServiceInterface) chi.Router {
	router := chi.NewRouter()

	listController := controllers.LogEntryListController{DbService: dbService}
	singleController := controllers.LogEntrySingleController{DbService: dbService}

	router.Get("/", listController.Get)
	router.Post("/", listController.Post)
	router.Route(fmt.Sprintf("/{%s}", controllers.PARAM_LOG_ENTRY_ID), func(router chi.Router) {
		router.Use(singleController.Context)
		router.Get("/", singleController.Get)
		router.Put("/", singleController.Put)
		router.Delete("/", singleController.Delete)
	})

	return router
}
