package log_file

import (
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
)

func TestParseLogFileLineWithSpaces(test *testing.T) {
	line := "    message    "
	dbService := &database.MockedDatabaseService{}
	logEntryDbService := log_entry.MockedLogEntryDatabaseService{DbService: dbService}

	service := LogDiffParserService{
		Producer:          "producer-name",
		Diff:              line,
		LogEntryDbService: logEntryDbService,
	}

	parsedLine, created, err := service.parseLogFileLine(line)
	if err != nil {
		test.Fatalf("An error has ocurred while parsing the line '%s': %s", line, err)
	}

	if parsedLine != "message" {
		test.Fatalf("Failed asserting that a line with leading and trailing spaces would be trimmed!")
	}

	if !created {
		test.Fatalf("Failed asserting that a line with valid text would generate a new log entry!")
	}

	test.Log("Finished testing the parseLogFileLine() method with a string containing leading and trailing spaces.")
}

func TestParseLogFileLineWithEmptyText(test *testing.T) {
	line := ""
	dbService := &database.MockedDatabaseService{}
	logEntryDbService := log_entry.MockedLogEntryDatabaseService{DbService: dbService}

	service := LogDiffParserService{
		Producer:          "producer-name",
		Diff:              line,
		LogEntryDbService: logEntryDbService,
	}

	parsedLine, created, err := service.parseLogFileLine(line)
	if err != nil {
		test.Fatalf("An error has ocurred while parsing the line '%s': %s", line, err)
	}

	if parsedLine != "" {
		test.Fatalf("Failed asserting that an empty line will be parsed as an empty string!")
	}

	if created {
		test.Fatalf("Failed asserting that an empty line would not create a new log entry!")
	}

	test.Log("Finished testing the parseLogFileLine() method with an empty string.")
}

func TestParseLogFileLineWithDuplicateText(test *testing.T) {
	line := "test"
	dbService := &database.MockedDatabaseService{}
	logEntryDbService := log_entry.MockedLogEntryDatabaseService{DbService: dbService}

	service := LogDiffParserService{
		Producer:          "producer-name",
		Diff:              line,
		LogEntryDbService: logEntryDbService,
	}

	parsedLine, created, err := service.parseLogFileLine(line)
	if err != nil {
		test.Fatalf("An error has ocurred while parsing the line '%s': %s", line, err)
	}

	if parsedLine != "test" {
		test.Fatalf("Failed asserting that a non-empty string would be parsed successfully!")
	}

	if created {
		test.Fatalf("Failed asserting that a line that already exists would not create a new entry!")
	}

	test.Log("Finished testing the parseLogFileLine() method with an empty string.")
}
