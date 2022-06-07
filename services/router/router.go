package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	chiRender "github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/services/authentication"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/docs"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/render"
)

func CreateRouter(dbService database.DatabaseServiceInterface, generateDocs bool) chi.Router {
	router := handler.NewHandler()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(chiRender.SetContentType(chiRender.ContentTypeJSON))

	renderService := &render.RenderService{}

	router.Mount("/health", HealthRouter(dbService, renderService))
	router.Mount("/log-entries", authentication.BasicAuthHandler(LogEntryRouter(dbService)))

	// Passing -generate-docs=false will skip the documentation generation
	if generateDocs {
		docs.ExportApiDocumentation(router)
	}

	return router
}
