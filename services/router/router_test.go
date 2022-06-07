package router

import (
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
)

func TestCreateRouter(test *testing.T) {
	dbService := &database.MockedDatabaseService{}
	router := CreateRouter(dbService, false)

	expectedRoutesCount := 2
	actualRoutesCount := len(router.Routes())
	if expectedRoutesCount != actualRoutesCount {
		test.Fatalf("Failed asserting that the router has the expected number of %d routes (atual: %d)!", expectedRoutesCount, actualRoutesCount)
	}

	routeZeroPattern := router.Routes()[0].Pattern
	if routeZeroPattern != "/health/*" {
		test.Fatalf("Failed asserting that the route #0 pattern is equal to '/health/*' (actual: '%s')!", routeZeroPattern)
	}

	routeOnePattern := router.Routes()[1].Pattern
	if routeOnePattern != "/log-entries/*" {
		test.Fatalf("Failed asserting that the route #1 pattern is equal to '/log-entries/*' (actual: '%s')!", routeOnePattern)
	}

	test.Log("Finished testing the CreateRouter() method")
}
