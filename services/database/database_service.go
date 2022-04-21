package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/environment"
	"github.com/ferdn4ndo/userver-logger-api/services/file"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("No matching record.")

const EMPTY_DB_FIXTURE = "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture/empty.sqlite.db"

type DatabaseService struct {
	Conn *gorm.DB
}

func (db *DatabaseService) Close() {
	sqlDB, err := db.Conn.DB()

	if err != nil {
		panic("Error closing DB connection!")
	}

	// Close
	sqlDB.Close()
}

func AddHeartbeatLog(db *DatabaseService) error {
	currentTime := time.Now().Format(time.RFC3339)
	result := db.Conn.Create(&models.LogEntry{
		Producer: "userver-logger-api",
		Message:  "Heartbeat from userver-logger-api at " + currentTime})

	return result.Error
}

func getDatabaseFilePath() string {
	dataFolder := file.GetDataFolder()
	databaseFile := environment.GetEnvKey("DATABASE_FILE")

	return fmt.Sprintf("%s/%s", dataFolder, databaseFile)
}

func geEmptyFixtureFilePath() string {
	fixtureFolder := file.GetFixtureFolder()
	emptyDatabaseFile := environment.GetEnvKey("EMPTY_DATABASE_FILE")

	return fmt.Sprintf("%s/%s", fixtureFolder, emptyDatabaseFile)
}

func createEmptyDatabase() error {
	log.Println("Creating empty database...")

	databasePath := getDatabaseFilePath()
	if _, err := os.Stat(databasePath); err == nil {
		log.Panicf("Database file '%s' already exists!", databasePath)
	}

	emptyFixturePath := geEmptyFixtureFilePath()
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

func InitializeDatabase() (*DatabaseService, error) {
	databasePath := getDatabaseFilePath()

	if _, err := os.Stat(databasePath); errors.Is(err, os.ErrNotExist) {
		err = createEmptyDatabase()
		if err != nil {
			return nil, err
		}
	}

	conn := sqlite.Open(databasePath)

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	service := &DatabaseService{Conn: db}

	log.Println("Migrating the schema...")
	service.Conn.AutoMigrate(&models.LogEntry{})

	AddHeartbeatLog(service)

	log.Println("Database connection established!")

	return service, nil
}

func GetDatabaseService() (*DatabaseService, error) {
	databasePath := getDatabaseFilePath()
	conn := sqlite.Open(databasePath)

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	service := &DatabaseService{Conn: db}

	return service, nil
}
