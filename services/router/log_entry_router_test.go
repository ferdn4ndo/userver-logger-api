package router

import (
	"testing"
)

func TestLogEntryRouter(test *testing.T) {
	router := LogEntryRouter()

	expectedRoutesCount := 2
	actualRoutesCount := len(router.Routes())
	if expectedRoutesCount != actualRoutesCount {
		test.Fatalf("Failed asserting that the router has the expected number of %d routes (atual: %d)!", expectedRoutesCount, actualRoutesCount)
	}

	routeZeroPattern := router.Routes()[0].Pattern
	if "/" != routeZeroPattern {
		test.Fatalf("Failed asserting that the route #0 pattern is equal to '/' (actual: '%s')!", routeZeroPattern)
	}

	routeOnePattern := router.Routes()[1].Pattern
	if "/{logEntryId}/*" != routeOnePattern {
		test.Fatalf("Failed asserting that the route #1 pattern is equal to '/{logEntryId}/*' (actual: '%s')!", routeOnePattern)
	}

	test.Log("Finished testing the LogEntryRouter() method")
}
