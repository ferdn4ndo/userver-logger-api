package database

import (
	"fmt"
	"testing"
)

func TestGetDatabaseFilePath(test *testing.T) {
	expectedDbPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/data/sqlite.db"
	computedDbPath := getDatabaseFilePath()

	if expectedDbPath != computedDbPath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed DB path '%s' is equal to expected '%s'.", computedDbPath, expectedDbPath))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}

func TestGetEmptyFixtureFilePath(test *testing.T) {
	expectedFixturePath := "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture/empty.sqlite.db"
	computedFixturePath := getEmptyFixtureFilePath()

	if expectedFixturePath != computedFixturePath {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed DB empty fixture path '%s' is equal to expected '%s'.", computedFixturePath, expectedFixturePath))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}
