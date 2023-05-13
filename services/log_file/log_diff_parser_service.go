package log_file

import (
	"bufio"
	"strings"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type LogDiffParserService struct {
	Producer          string
	Diff              string
	LogEntryDbService log_entry.LogEntryDatabaseServiceInterface
}

func (service LogDiffParserService) ParseDiff() error {
	scanner := bufio.NewScanner(strings.NewReader(service.Diff))
	for scanner.Scan() {
		line := scanner.Text()

		_, _, err := service.parseLogFileLine(line)
		if err != nil {
			logging.Errorf("Error parsing line '%s': %s", line, err)
		}
	}

	if err := scanner.Err(); err != nil {
		logging.Errorf("Error scanning string: %s", err)
	}

	return nil
}

func (service LogDiffParserService) parseLogFileLine(line string) (string, bool, error) {
	line = strings.TrimSpace(line)

	if line == "" {
		logging.Debugf("Skipping empty diff line!")

		return line, false, nil
	}

	logEntryExists, err := service.LogEntryDbService.CheckIfLogEntryExists(service.Producer, line)
	if err != nil {
		logging.Errorf("Error checking if log entry line '%s' exists: %s", line, err)

		return line, false, err
	}

	if !logEntryExists {
		model := &models.LogEntry{
			Producer: service.Producer,
			Message:  line,
		}

		if err := service.LogEntryDbService.AddLogEntry(model); err != nil {
			logging.Errorf("Error adding parsed log entry for line '%s': %s", line, err)

			return line, false, err
		}

		return line, true, nil
	} else {
		logging.Debugf("Skipping duplicate line: %s", line)

		return line, false, nil
	}
}
