package router

import (
	"testing"
)

func TestHealthRouter(test *testing.T) {
	router := HealthRouter()

	expectedRoutesCount := 1
	actualRoutesCount := len(router.Routes())
	if expectedRoutesCount != actualRoutesCount {
		test.Fatalf("Failed asserting that the router has the expected number of %d routes (atual: %d)!", expectedRoutesCount, actualRoutesCount)
	}

	routeZeroPattern := router.Routes()[0].Pattern
	if "/" != routeZeroPattern {
		test.Fatalf("Failed asserting that the route #0 pattern is equal to '/' (actual: '%s')!", routeZeroPattern)
	}

	test.Log("Finished testing the HealthRouter() method")
}
