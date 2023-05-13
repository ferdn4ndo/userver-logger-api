package middleware

import (
	"testing"
)

func TestGetCorsHeaderValueWithEmptyValue(test *testing.T) {
	corsValue := getCorsHeaderValue("")

	if corsValue != "*" {
		test.Fatal("Failed asserting that the CORS value will be the '*' wildcard when an empty argument is provided!")
	}

	test.Log("Finished testing the getCorsHeaderValue() method with an empty value!")
}

func TestGetCorsHeaderValueWithNonEmptyValue(test *testing.T) {
	expectedValue := "domain.lan"
	corsValue := getCorsHeaderValue(expectedValue)

	if corsValue != expectedValue {
		test.Fatal("Failed asserting that the CORS value will be the same as the provided argument!")
	}

	test.Log("Finished testing the getCorsHeaderValue() method with a non-empty value!")
}
