package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/environment"
	"github.com/ferdn4ndo/userver-logger-api/services/file"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

const EMPTY_DB_FIXTURE = "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture/empty.sqlite.db"

type DatabaseServiceInterface interface {
	GetDbConn() *gorm.DB
	AddHeartbeatLog() error
	Close()
	GetDatabaseFileSize() int64
	GetLogEntriesTotalCount() int64
}

type DatabaseService struct {
	conn *gorm.DB
}

func (db DatabaseService) Close() {
	sqlDB, err := db.GetDbConn().DB()

	if err != nil {
		panic("Error closing DB connection!")
	}

	// Close
	sqlDB.Close()
}

func (db DatabaseService) GetDbConn() *gorm.DB {
	return db.conn
}

func (db DatabaseService) GetDatabaseFileSize() int64 {
	file, err := os.Open(getDatabaseFilePath())
	if err != nil {
		logging.Errorf("Error opening database file: %s", err)
	}

	fileStats, err := file.Stat()
	if err != nil {
		logging.Errorf("Error checking database file: %s", err)
	}

	return fileStats.Size()
}

func (db DatabaseService) AddHeartbeatLog() error {
	currentTime := time.Now().Format(time.RFC3339)
	result := db.GetDbConn().Create(&models.LogEntry{
		Producer: "userver-logger-api",
		Message:  "Heartbeat from userver-logger-api at " + currentTime})

	return result.Error
}

func (db DatabaseService) GetLogEntriesTotalCount() int64 {
	var logEntriesCount int64
	db.GetDbConn().Model(&models.LogEntry{}).Count(&logEntriesCount)

	return logEntriesCount
}

func getDatabaseFilePath() string {
	dataFolder := file.GetDataFolder()
	databaseFile := environment.GetEnvKey("DATABASE_FILE")

	return fmt.Sprintf("%s/%s", dataFolder, databaseFile)
}

func getEmptyFixtureFilePath() string {
	fixtureFolder := file.GetFixtureFolder()
	emptyDatabaseFile := environment.GetEnvKey("EMPTY_DATABASE_FILE")

	return fmt.Sprintf("%s/%s", fixtureFolder, emptyDatabaseFile)
}

func createEmptyDatabase() error {
	logging.Info("Creating empty database...")

	databasePath := getDatabaseFilePath()
	if _, err := os.Stat(databasePath); err == nil {
		logging.Errorf("Database file '%s' already exists!", databasePath)
	}

	emptyFixturePath := getEmptyFixtureFilePath()
	emptyFixtureData, err := ioutil.ReadFile(emptyFixturePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(databasePath, emptyFixtureData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func InitializeDatabase() DatabaseServiceInterface {
	databasePath := getDatabaseFilePath()

	if _, err := os.Stat(databasePath); errors.Is(err, os.ErrNotExist) {
		err = createEmptyDatabase()
		if err != nil {
			logging.Errorf("Error creating empty database: %s", err)
		}
	}

	conn := sqlite.Open(databasePath)

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		logging.Errorf("Error opening DB connection: %s", err)
	}

	service := DatabaseService{conn: db}

	logging.Debug("Migrating the schema...")
	if err = service.conn.AutoMigrate(&models.LogEntry{}); err != nil {
		logging.Errorf("Error applying DB migrations: %s", err)
	}

	if err := service.AddHeartbeatLog(); err != nil {
		logging.Errorf("Error adding heartbeat: %s", err)
	}

	logging.Debug("Database connection established!")

	return service
}

func GetDatabaseService() (DatabaseService, error) {
	databasePath := getDatabaseFilePath()
	conn := sqlite.Open(databasePath)

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		return DatabaseService{}, err
	}

	service := DatabaseService{conn: db}

	return service, nil
}
