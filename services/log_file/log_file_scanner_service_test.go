package log_file

import (
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/file"
)

func TestFindLogFiles(test *testing.T) {
	logFiles := findLogFiles(file.GetLogFilesFolder(), ".log")
	logFilesCount := len(logFiles)

	if logFilesCount != 1 {
		test.Fatalf("Failed asserting that the fixture folder log file list has one entry (actual count: %d).", logFilesCount)
	}

	test.Log("Finished testing the findLogFiles() method")
}
