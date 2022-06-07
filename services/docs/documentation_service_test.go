package docs

import (
	"fmt"
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

func TestGetApiDocsJsonFile(test *testing.T) {
	expectedDocsPath := fmt.Sprintf("%s/routes.json", environment.GetEnvKey("DATA_FOLDER"))
	computedDocsPath := getApiDocsJsonFile()

	if expectedDocsPath != computedDocsPath {
		test.Fatalf("Failed asserting that the computed API Docs json path '%s' is equal to expected '%s'.", computedDocsPath, expectedDocsPath)
	}

	test.Log("Finished testing the getApiDocsJsonFile() method")
}
