package router

import (
	"flag"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/services/authentication"
	"github.com/ferdn4ndo/userver-logger-api/services/docs"
)

var generateDocs = flag.Bool("generate-docs", true, "Generate router documentation")

func CreateRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Mount("/health", HealthRouter())
	router.Mount("/log-entries", authentication.BasicAuthHandler(LogEntryRouter()))

	// Passing -generate-docs=false will skip the documentation generation
	if *generateDocs {
		docs.ExportApiDocumentation(router)
	}

	return router
}
