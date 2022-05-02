package handler

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorRenderer(test *testing.T) {
	expectedMessage := "test error message"
	err := errors.New(expectedMessage)

	rendererPointer := ErrorRenderer(err)
	renderer := *rendererPointer

	if renderer.StatusCode != 400 {
		test.Fatalf(fmt.Sprintf("Failed asserting that the expected renderer status code is 400. (Actual: %d)", renderer.StatusCode))
	}

	if BadRequestStatusText != renderer.StatusText {
		test.Fatalf(fmt.Sprintf("Failed asserting that the expected renderer status message '%s' is equal to '%s'.", BadRequestStatusText, renderer.StatusText))
	}

	if expectedMessage != renderer.Message {
		test.Fatalf(fmt.Sprintf("Failed asserting that the expected renderer message '%s' is equal to '%s'.", expectedMessage, renderer.Message))
	}

	test.Log("Finished testing the ErrorRenderer() method")
}
