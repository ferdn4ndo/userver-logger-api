package log_entry

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/pagination"
)

type LogEntryServiceInterface interface {
	GetAllLogEntries(limit int, offset int, params *LogEntrySearchParams) (*models.LogEntryList, int, error)
	AddLogEntry(logEntry *models.LogEntry) error
	GetLogEntryById(logEntryId uint) (*models.LogEntry, error)
	UpdateLogEntry(logEntryId uint, logEntryData *models.LogEntryRequest) (*models.LogEntry, error)
	DeleteLogEntry(logEntryId uint) error
	CheckIfIdExists(logEntryId uint) (bool, error)
	CheckIfLogEntryExists(producer string, message string) (bool, error)
}

type LogEntryDatabaseService struct {
	DbService database.DatabaseServiceInterface
}

func (service LogEntryDatabaseService) GetAllLogEntries(request *http.Request, params *LogEntrySearchParams) (*models.LogEntryList, int, error) {
	var logEntries models.LogEntryList

	query := service.DbService.GetDbConn().Model(&models.LogEntry{})
	query = ApplyLogEntryQuerySearchParams(query, params)

	var totalCount int64
	query.Count(&totalCount)

	paginationService := &pagination.PaginationService{Request: request}
	paginationService.ApplyQueryOffsetAndLimit(query)

	listQuery := query.Find(&logEntries.LogEntries)
	if listQuery.Error != nil {
		return &logEntries, 0, listQuery.Error
	}

	return &logEntries, int(totalCount), nil
}

func (service LogEntryDatabaseService) AddLogEntry(logEntry *models.LogEntry) error {
	result := service.DbService.GetDbConn().Create(&logEntry)

	return result.Error
}

func (service LogEntryDatabaseService) CheckIfIdExists(logEntryId uint) (bool, error) {
	var exists bool

	err := service.
		DbService.
		GetDbConn().
		Model(&models.LogEntry{}).
		Select("count(*) > 0").
		Where("id = ?", logEntryId).
		Find(&exists).
		Error

	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (service LogEntryDatabaseService) GetLogEntryById(logEntryId uint) (*models.LogEntry, error) {
	var logEntry models.LogEntry

	error := service.DbService.GetDbConn().First(&logEntry, "id = ?", logEntryId).Error

	if error != nil {
		if error == gorm.ErrRecordNotFound {
			return &logEntry, database.ErrNoMatch
		} else {
			return &logEntry, error
		}
	}

	return &logEntry, nil
}

func (service LogEntryDatabaseService) DeleteLogEntry(logEntryId uint) error {
	var logEntry models.LogEntry

	error := service.DbService.GetDbConn().Delete(&logEntry, logEntryId).Error

	return error
}

func (service LogEntryDatabaseService) UpdateLogEntry(logEntryId uint, logEntryData *models.LogEntryRequest) (*models.LogEntry, error) {
	var logEntry models.LogEntry

	error := service.DbService.GetDbConn().First(&logEntry, "id = ?", logEntryId).Error
	if error != nil {
		return &logEntry, error
	}

	error = service.DbService.GetDbConn().Model(&logEntry).Updates(logEntryData.LogEntry).Error
	if error != nil {
		return &logEntry, error
	}

	return &logEntry, nil
}

func (service LogEntryDatabaseService) CheckIfLogEntryExists(producer string, message string) (bool, error) {
	var logEntry models.LogEntry
	var count int64

	service.DbService.GetDbConn().Model(&logEntry).Where("producer = ?", producer).Where("message = ?", message).Count(&count)

	return count > 0, nil
}
