package controllers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/http_request"
	"github.com/ferdn4ndo/userver-logger-api/services/render"
)

func TestGetHealthState(test *testing.T) {
	controller := HealthController{DbService: &database.MockedDatabaseService{}, RenderService: render.MockedRenderService{}}

	mockedWriter := http_request.MockedHttpResponseWriter{}
	mockedRequest, err := http.NewRequest(http.MethodGet, "/health", bytes.NewBufferString(""))
	if err != nil {
		log.Fatalf("Error mocking request: %s", err)
	}

	rescueStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	controller.GetHealthState(mockedWriter, mockedRequest)

	writer.Close()
	out, _ := ioutil.ReadAll(reader)
	os.Stdout = rescueStdout

	expectedOutput := "{\"status\":\"OK\",\"databaseSizeInBytes\":1024,\"logEntriesCount\":100}"
	if string(out) != expectedOutput {
		test.Errorf("Expected %s, got %s", expectedOutput, out)
	}

	test.Log("Finished testing the GetHealthState() method")
}
