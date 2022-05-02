package log_file

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
)

func getContainerNameFromPath(fullFilePath string) string {
	filename := filepath.Base(fullFilePath)
	extension := filepath.Ext(filename)

	return strings.TrimSuffix(filename, extension)
}

func parseLogFileLine(containerName string, line string) {
	logEntryExists, err := log_entry.CheckIfLogEntryExists(containerName, line)
	if err != nil {
		log.Fatalf("Error checking if log entry exists: %s", err)
	}

	if !logEntryExists {
		model := &models.LogEntry{
			Producer: containerName,
			Message:  line,
		}

		if err := log_entry.AddLogEntry(model); err != nil {
			log.Fatalf("Error adding parsed log line entry: %s", err)
		}
	}
}

func ParseLogFile(filename string) error {
	containerName := getContainerNameFromPath(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening log file '%s': %s", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseLogFileLine(containerName, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning log file '%s': %s", filename, err)
	}

	return nil
}
