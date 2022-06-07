package log_file

import (
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/ferdn4ndo/userver-logger-api/services/checksum"
	"github.com/ferdn4ndo/userver-logger-api/services/file"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type LogFileScannerService struct {
	LogEntryDbService log_entry.LogEntryDatabaseService
}

func (service LogFileScannerService) ForeverScanLogFiles() {
	logFileScannerService := &LogFileScannerService{LogEntryDbService: service.LogEntryDbService}

	logging.Debug("Scanning log files forever...")
	for {
		logFileScannerService.ScanLogFiles()
		time.Sleep(time.Second * 5)
	}
}

func (service LogFileScannerService) ScanLogFiles() {
	logsFolderPath := file.GetLogFilesFolder()
	logFiles := findLogFiles(logsFolderPath, ".log")

	for _, file := range logFiles {
		service.CheckLogFile(file)
	}
}

func (service LogFileScannerService) CheckLogFile(filename string) {
	cachedFileChecksum := checksum.GetCachedFileChecksum(filename)

	computedFileChecksum, err := checksum.ComputeFileChecksum(filename)
	if err != nil {
		logging.Errorf("Error checking log file '%s': %s", filename, err)
	}

	needsChecksumUpdate := cachedFileChecksum == "" || cachedFileChecksum != computedFileChecksum
	if !needsChecksumUpdate {
		logging.Debugf("File '%s' SHA256 was not updated, skipping...", filename)

		return
	}

	logging.Debugf("SHA256 checksum for file '%s' has changed!", filename)
	checksum.SetCachedFileChecksum(filename, computedFileChecksum)

	logging.Debugf("Checking if update is required for filename '%s'", filename)
	consumedFileService := &file.ConsumedLinesFileService{LogFilePath: filename}
	if !consumedFileService.RequiresUpdate() {
		logging.Debugf("No differences detected between current log file and the last consumed one, skipping...")

		return
	}

	logging.Infof("Detected diff for last consumed lines on file '%s', parsing...", filename)
	diff, err := consumedFileService.GetLastConsumedFileDiff()
	if err != nil {
		logging.Errorf("Error getting last consumed file diff: %s", err)
	}

	diffParserService := &LogDiffParserService{
		Producer:          consumedFileService.GetProducer(),
		Diff:              diff,
		LogEntryDbService: service.LogEntryDbService,
	}

	logging.Debugf("Parsing file diff...")
	if err := diffParserService.ParseDiff(); err != nil {
		logging.Errorf("Error parsing diff: %s", err)
	}

	logging.Debugf("Updating last consumed file...")
	if err := consumedFileService.UpdateLastConsumedLines(); err != nil {
		logging.Errorf("Error updating last consumed file: %s", err)
	}
}

func findLogFiles(root, ext string) []string {
	var filesList []string

	err := filepath.WalkDir(root, func(filename string, fileEntry fs.DirEntry, e error) error {
		if e != nil {
			logging.Errorf("Error reading log file '%s' inside folder '%s': %s", filename, root, e)
		}

		if strings.EqualFold(filepath.Ext(fileEntry.Name()), ext) {
			filesList = append(filesList, filename)
		}

		return nil
	})

	if err != nil {
		logging.Errorf("Error finding log files inside folder '%s': %s", root, err)
	}

	return filesList
}
