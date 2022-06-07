package router

import (
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
)

func TestLogEntryRouter(test *testing.T) {
	dbService := &database.MockedDatabaseService{}
	router := LogEntryRouter(dbService)

	expectedRoutesCount := 2
	actualRoutesCount := len(router.Routes())
	if expectedRoutesCount != actualRoutesCount {
		test.Fatalf("Failed asserting that the router has the expected number of %d routes (atual: %d)!", expectedRoutesCount, actualRoutesCount)
	}

	routeZeroPattern := router.Routes()[0].Pattern
	if routeZeroPattern != "/" {
		test.Fatalf("Failed asserting that the route #0 pattern is equal to '/' (actual: '%s')!", routeZeroPattern)
	}

	routeOnePattern := router.Routes()[1].Pattern
	if routeOnePattern != "/{logEntryId}/*" {
		test.Fatalf("Failed asserting that the route #1 pattern is equal to '/{logEntryId}/*' (actual: '%s')!", routeOnePattern)
	}

	test.Log("Finished testing the LogEntryRouter() method")
}
