package router

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	chiRender "github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/services/authentication"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/docs"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/middleware"
	"github.com/ferdn4ndo/userver-logger-api/services/render"
)

func CreateRouter(dbService database.DatabaseServiceInterface, generateDocs bool) chi.Router {
	router := handler.NewHandler()

	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.URLFormat)

	router.Use(middleware.AddCorsHeader)

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
