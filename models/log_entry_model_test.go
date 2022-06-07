package models

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestBind(test *testing.T) {
	model := LogEntry{}

	mockedRequest, err := http.NewRequest(http.MethodGet, "/", bytes.NewBufferString(""))
	if err != nil {
		log.Fatalf("Error mocking request: %s", err)
	}

	err = model.Bind(mockedRequest)
	if err.Error() != "the field 'producer' is required" {
		test.Fatalf("Failed asserting that an error is raised when the producer field is not set!")
	}

	model.Producer = "test"
	err = model.Bind(mockedRequest)
	if err.Error() != "the field 'message' is required" {
		test.Fatalf("Failed asserting that an error is raised when the message field is not set!")
	}

	model.Message = "test"
	err = model.Bind(mockedRequest)
	if err != nil {
		test.Fatalf("Failed asserting that no error is raised when both the producer and message fields are given!")
	}

	test.Log("Finished testing the Bind() method")
}
