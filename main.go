package main

import (
	"flag"

	"github.com/ferdn4ndo/userver-logger-api/services/application"
)

var (
	generateDocs = flag.Bool("generate-docs", true, "Generate router documentation")
	dryRun       = flag.Bool("dry-run", false, "Don't connect to real database")
)

func main() {
	applicationStartService := &application.ApplicationStartService{
		GenerateDocs: *generateDocs,
		DryRun:       *dryRun,
	}
	applicationStartService.StartApplication()
}
