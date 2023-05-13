package log_file

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type LogDiffParserService struct {
	Producer          string
	Diff              string
	LogEntryDbService log_entry.LogEntryDatabaseService
}

func (service LogDiffParserService) ParseDiff() error {
	scanner := bufio.NewScanner(strings.NewReader(service.Diff))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		service.parseLogFileLine(line)
	}

	if err := scanner.Err(); err != nil {
		logging.Errorf("Error scanning string: %s", err)
	}

	return nil
}

func (service LogDiffParserService) parseLogFileLine(line string) {
	logEntryExists, err := service.LogEntryDbService.CheckIfLogEntryExists(service.Producer, line)
	if err != nil {
		logging.Error(fmt.Sprintf("Error checking if log entry exists: %s", err))
	}

	if !logEntryExists {
		model := &models.LogEntry{
			Producer: service.Producer,
			Message:  line,
		}

		if err := service.LogEntryDbService.AddLogEntry(model); err != nil {
			logging.Errorf("Error adding parsed log line entry: %s", err)
		}
	} else {
		logging.Debug(fmt.Sprintf("Skipping duplicate line: %s", line))
	}
}
