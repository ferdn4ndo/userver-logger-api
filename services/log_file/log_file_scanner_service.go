package log_file

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/ferdn4ndo/userver-logger-api/services/checksum"
	"github.com/ferdn4ndo/userver-logger-api/services/file"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
)

type LogFileScannerService struct {
	LogEntryDbService log_entry.LogEntryDatabaseService
}

func (service LogFileScannerService) ForeverScanLogFiles() {
	logFileScannerService := &LogFileScannerService{LogEntryDbService: service.LogEntryDbService}

	log.Printf("Starting scanning log files forever...")
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
		log.Panicf("Error checking log file '%s': %s", filename, err)
	}

	needsChecksumUpdate := cachedFileChecksum == "" || cachedFileChecksum != computedFileChecksum
	if !needsChecksumUpdate {
		// log.Printf("File '%s' SHA256 was not updated, skipping...", filename)
		return
	}

	log.Printf("SHA256 checksum for file '%s' has changed!", filename)
	checksum.SetCachedFileChecksum(filename, computedFileChecksum)

	consumedFileService := &file.ConsumedLinesFileService{LogFilePath: filename}
	if !consumedFileService.RequiresUpdate() {
		log.Print("No differences detected between current log file and the last consumed one, skipping...")
		return
	}

	log.Printf("Detected diff for last consumed lines on file '%s', updating...", filename)

	diff, err := consumedFileService.GetLastConsumedFileDiff()
	if err != nil {
		log.Fatalf("Error getting last consumed file diff: %s", err)
	}

	diffParserService := &LogDiffParserService{
		Producer:          consumedFileService.GetProducer(),
		Diff:              diff,
		LogEntryDbService: service.LogEntryDbService,
	}

	err = diffParserService.ParseDiff()
	if err != nil {
		log.Fatalf("Error parsing diff: %s", err)
	}
}

func findLogFiles(root, ext string) []string {
	var filesList []string

	err := filepath.WalkDir(root, func(filename string, fileEntry fs.DirEntry, e error) error {
		if e != nil {
			log.Panicf("Error reading log file '%s' inside folder '%s': %s", filename, root, e)
		}

		if strings.EqualFold(filepath.Ext(fileEntry.Name()), ext) {
			filesList = append(filesList, filename)
		}

		return nil
	})

	if err != nil {
		log.Panicf("Error finding log files inside folder '%s': %s", root, err)
	}

	return filesList
}
