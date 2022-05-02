package file

import (
	"fmt"
	"testing"
)

func TestGetTempFolder(test *testing.T) {
	expectedTempPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/tmp"
	computedTempPath := GetTempFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath))
	}

	test.Log("Finished testing the GetTempFolder() method")
}

func TestGetDataFolder(test *testing.T) {
	expectedTempPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/data"
	computedTempPath := GetDataFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath))
	}

	test.Log("Finished testing the GetDataFolder() method")
}

func TestFixtureFolder(test *testing.T) {
	expectedTempPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture"
	computedTempPath := GetFixtureFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath))
	}

	test.Log("Finished testing the GetFixtureFolder() method")
}

func TestLogFilesFolder(test *testing.T) {
	expectedTempPath := "/log_files"
	computedTempPath := GetLogFilesFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath))
	}

	test.Log("Finished testing the GetLogFilesFolder() method")
}
