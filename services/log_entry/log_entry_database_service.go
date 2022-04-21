package log_entry

import (
	"gorm.io/gorm"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/pagination"
)

func GetAllLogEntries(limit int, offset int, params *LogEntrySearchParams) (*models.LogEntryList, int, error) {
	var logEntries models.LogEntryList

	db, error := database.GetDatabaseService()
	if error != nil {
		return &logEntries, 0, error
	}

	query := db.Conn.Model(&models.LogEntry{})
	query = ApplyLogEntryQuerySearchParams(query, params)

	var totalCount int64
	query.Count(&totalCount)

	query = pagination.ApplyQueryOffsetAndLimit(query, offset, limit)

	list := query.Find(&logEntries.LogEntries)
	db.Close()

	if list.Error != nil {
		return &logEntries, 0, list.Error
	}

	return &logEntries, int(totalCount), nil
}

func AddLogEntry(logEntry *models.LogEntry) error {
	db, error := database.GetDatabaseService()
	if error != nil {
		return error
	}

	result := db.Conn.Create(&logEntry)
	db.Close()

	return result.Error
}

func CheckIfIdExists(logEntryId uint) (bool, error) {
	var exists bool

	db, error := database.GetDatabaseService()
	if error != nil {
		return exists, error
	}

	error = db.Conn.Model(&models.LogEntry{}).
		Select("count(*) > 0").
		Where("id = ?", logEntryId).
		Find(&exists).Error
	db.Close()

	if error != nil {
		return exists, error
	}

	return exists, nil
}

func GetLogEntryById(logEntryId uint) (*models.LogEntry, error) {
	var logEntry models.LogEntry

	db, error := database.GetDatabaseService()
	if error != nil {
		return &logEntry, error
	}

	error = db.Conn.First(&logEntry, "id = ?", logEntryId).Error
	db.Close()

	if error != nil {
		if error == gorm.ErrRecordNotFound {
			return &logEntry, database.ErrNoMatch
		} else {
			return &logEntry, error
		}
	}

	return &logEntry, nil
}

func DeleteLogEntry(logEntryId uint) error {
	var logEntry models.LogEntry

	db, error := database.GetDatabaseService()
	if error != nil {
		return error
	}

	error = db.Conn.Delete(&logEntry, logEntryId).Error
	db.Close()

	return error
}

func UpdateLogEntry(logEntryId uint, logEntryData *models.LogEntryRequest) (*models.LogEntry, error) {
	var logEntry models.LogEntry

	db, error := database.GetDatabaseService()
	if error != nil {
		return &logEntry, error
	}

	error = db.Conn.First(&logEntry, "id = ?", logEntryId).Error
	if error != nil {
		return &logEntry, error
	}

	error = db.Conn.Model(&logEntry).Updates(logEntryData.LogEntry).Error
	if error != nil {
		return &logEntry, error
	}

	db.Close()

	return &logEntry, nil
}

func CheckIfLogEntryExists(producer string, message string) (bool, error) {
	var logEntry models.LogEntry
	var count int64

	db, error := database.GetDatabaseService()
	if error != nil {
		return false, error
	}

	db.Conn.Model(&logEntry).Where("producer = ?", producer).Where("message = ?", message).Count(&count)

	return count > 0, nil
}
