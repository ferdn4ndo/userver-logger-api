package log_entry

import (
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
)

type MockedLogEntryDatabaseService struct {
	DbService database.DatabaseServiceInterface
}

func (service MockedLogEntryDatabaseService) GetAllLogEntries(request *http.Request, params *LogEntrySearchParams) (*models.LogEntryList, int, error) {
	var logEntries models.LogEntryList

	return &logEntries, 0, nil
}

func (service MockedLogEntryDatabaseService) AddLogEntry(logEntry *models.LogEntry) error {
	return nil
}

func (service MockedLogEntryDatabaseService) CheckIfIdExists(logEntryId uint) (bool, error) {
	if logEntryId == 420 {
		return true, nil
	}

	return false, nil
}

func (service MockedLogEntryDatabaseService) GetLogEntryById(logEntryId uint) (*models.LogEntry, error) {
	var logEntry models.LogEntry

	if logEntryId == 420 {
		logEntry.Message = "test"
		logEntry.ID = 420
		logEntry.Producer = "test"

		return &logEntry, nil
	}

	return &logEntry, database.ErrNoMatch
}

func (service MockedLogEntryDatabaseService) DeleteLogEntry(logEntryId uint) error {
	return nil
}

func (service MockedLogEntryDatabaseService) UpdateLogEntry(logEntryId uint, logEntryData *models.LogEntryRequest) (*models.LogEntry, error) {
	var logEntry models.LogEntry

	return &logEntry, nil
}

func (service MockedLogEntryDatabaseService) CheckIfLogEntryExists(producer string, message string) (bool, error) {
	if message == "test" {
		return true, nil
	}

	return false, nil
}
