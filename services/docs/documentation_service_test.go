package docs

import (
	"fmt"
	"testing"
)

func TestGetApiDocsJsonFile(test *testing.T) {
	expectedDocsPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/data/routes.json"
	computedDocsPath := getApiDocsJsonFile()

	if expectedDocsPath != computedDocsPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed API Docs json path '%s' is equal to expected '%s'.", computedDocsPath, expectedDocsPath))
	}

	test.Log("Finished testing the getApiDocsJsonFile() method")
}
