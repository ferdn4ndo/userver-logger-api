package database

import (
	"gorm.io/gorm"

	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type MockedDatabaseService struct{}

func (MockedDatabaseService) GetDbConn() *gorm.DB {
	return nil
}

func (MockedDatabaseService) AddHeartbeatLog() error {
	return nil
}

func (MockedDatabaseService) Close() {
	logging.Debug("Mocking Close call")
}

func (MockedDatabaseService) GetDatabaseFileSize() int64 {
	return 1024
}

func (MockedDatabaseService) GetLogEntriesTotalCount() int64 {
	return 100
}
