package health

import (
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
)

func TestGetHealthData(test *testing.T) {
	dbService := &database.MockedDatabaseService{}
	healthService := HealthService{DbService: dbService}
	healthData := healthService.GetHealthData()

	if healthData.Status != HEALTH_STATUS_OK {
		test.Fatalf("Failed asserting that health status output is %s (actual: '%s')!", HEALTH_STATUS_OK, healthData.Status)
	}

	expectedDbSize := dbService.GetDatabaseFileSize()
	actualDbSize := healthData.DbSize
	if expectedDbSize != actualDbSize {
		test.Fatalf("Failed asserting that the database size %d matched the expected %d!", actualDbSize, expectedDbSize)
	}

	expectedLogCount := dbService.GetLogEntriesTotalCount()
	actualLogCount := healthData.LogEntriesCount
	if expectedLogCount != actualLogCount {
		test.Fatalf("Failed asserting that the total logs count %d matched the expected %d!", actualLogCount, expectedLogCount)
	}

	test.Log("Finished testing the GetHealthData() method")
}
