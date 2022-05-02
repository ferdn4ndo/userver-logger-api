package log_file

import (
	"fmt"
	"os"
	"testing"
)

func TestFindLogFiles(test *testing.T) {
	logFiles := findLogFiles(os.Getenv("FIXTURE_FOLDER"), ".log")
	logFilesCount := len(logFiles)

	if 1 != logFilesCount {
		test.Fatalf(fmt.Sprintf("Failed asserting that the fixture folder log file list has one entry (actual count: %d).", logFilesCount))
	}

	test.Log("Finished testing the findLogFiles() method")
}
