package database

import (
	"fmt"
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

func TestGetDatabaseFilePath(test *testing.T) {
	expectedDbPath := fmt.Sprintf("%s/sqlite.db", environment.GetEnvKey("DATA_FOLDER"))
	computedDbPath := getDatabaseFilePath()

	if expectedDbPath != computedDbPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed DB path '%s' is equal to expected '%s'.", computedDbPath, expectedDbPath))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}

func TestGetEmptyFixtureFilePath(test *testing.T) {
	expectedFixturePath := fmt.Sprintf("%s/empty.sqlite.db", environment.GetEnvKey("FIXTURE_FOLDER"))
	computedFixturePath := getEmptyFixtureFilePath()

	if expectedFixturePath != computedFixturePath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed DB empty fixture path '%s' is equal to expected '%s'.", computedFixturePath, expectedFixturePath))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}
