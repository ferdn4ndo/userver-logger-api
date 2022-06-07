package log_entry

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetContainerNameFromPath(test *testing.T) {
	expectedProducer := "test_producer"
	expectedMessage := "test_message"

	request, _ := http.NewRequest("GET", fmt.Sprintf("/?producer=%s&message=%s", expectedProducer, expectedMessage), nil)
	searchParams := GetLogEntrySearchParams(request)

	actualProducer := searchParams.Producer
	if expectedProducer != actualProducer {
		test.Fatalf(fmt.Sprintf("Failed asserting that the producer query string is '%s' (actual: '%s').", expectedProducer, actualProducer))
	}

	actualMessage := searchParams.Message
	if expectedMessage != actualMessage {
		test.Fatalf(fmt.Sprintf("Failed asserting that the producer query string is '%s' (actual: '%s').", expectedProducer, actualMessage))
	}

	test.Log("Finished testing the GetLogEntrySearchParams() method")
}
