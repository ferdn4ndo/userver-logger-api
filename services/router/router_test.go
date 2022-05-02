package router

import (
	"testing"
)

func TestCreateRouter(test *testing.T) {
	router := CreateRouter()

	expectedRoutesCount := 2
	actualRoutesCount := len(router.Routes())
	if expectedRoutesCount != actualRoutesCount {
		test.Fatalf("Failed asserting that the router has the expected number of %d routes (atual: %d)!", expectedRoutesCount, actualRoutesCount)
	}

	routeZeroPattern := router.Routes()[0].Pattern
	if "/health/*" != routeZeroPattern {
		test.Fatalf("Failed asserting that the route #0 pattern is equal to '/health/*' (actual: '%s')!", routeZeroPattern)
	}

	routeOnePattern := router.Routes()[1].Pattern
	if "/log-entries/*" != routeOnePattern {
		test.Fatalf("Failed asserting that the route #1 pattern is equal to '/log-entries/*' (actual: '%s')!", routeOnePattern)
	}

	test.Log("Finished testing the CreateRouter() method")
}
