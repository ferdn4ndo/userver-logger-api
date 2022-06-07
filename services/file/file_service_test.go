package file

import (
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

func TestGetTempFolder(test *testing.T) {
	expectedTempPath := environment.GetEnvKey("TMP_FOLDER")
	computedTempPath := GetTempFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath)
	}

	test.Log("Finished testing the GetTempFolder() method")
}

func TestGetDataFolder(test *testing.T) {
	expectedTempPath := environment.GetEnvKey("DATA_FOLDER")
	computedTempPath := GetDataFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath)
	}

	test.Log("Finished testing the GetDataFolder() method")
}

func TestFixtureFolder(test *testing.T) {
	expectedTempPath := environment.GetEnvKey("FIXTURE_FOLDER")
	computedTempPath := GetFixtureFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath)
	}

	test.Log("Finished testing the GetFixtureFolder() method")
}

func TestLogFilesFolder(test *testing.T) {
	expectedTempPath := environment.GetEnvKey("LOG_FILES_FOLDER")
	computedTempPath := GetLogFilesFolder()

	if expectedTempPath != computedTempPath {
		test.Fatalf("Failed asserting that the computed temp path '%s' is equal to the expected '%s'.", computedTempPath, expectedTempPath)
	}

	test.Log("Finished testing the GetLogFilesFolder() method")
}

func TestGetContainerNameFromPath(test *testing.T) {
	filePath := "/test/path/container-name.log"
	expectedName := "container-name"
	computedName := GetContainerNameFromPath(filePath)

	if expectedName != computedName {
		test.Fatalf("Failed asserting that the computed container name '%s' is equal to the expected '%s'.", computedName, expectedName)
	}

	test.Log("Finished testing the GetContainerNameFromPath() method")
}
