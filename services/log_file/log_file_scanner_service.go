package log_file

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/ferdn4ndo/userver-logger-api/services/checksum"
	"github.com/ferdn4ndo/userver-logger-api/services/file"
)

func ForeverScanLogFiles() {
	log.Printf("Starting scanning log files forever...")
	for true {
		ScanLogFiles()
		time.Sleep(time.Second * 5)
	}
}

func ScanLogFiles() {
	logsFolderPath := file.GetLogFilesFolder()
	logFiles := findLogFiles(logsFolderPath, ".log")

	for _, file := range logFiles {
		CheckLogFile(file)
	}
}

func CheckLogFile(filename string) {
	cachedFileChecksum := checksum.GetCachedFileChecksum(filename)

	computedFileChecksum, err := checksum.ComputeFileChecksum(filename)
	if err != nil {
		log.Panicf("Error checking log file '%s': %s", filename, err)
	}

	needsUpdate := cachedFileChecksum == "" || cachedFileChecksum != computedFileChecksum
	if needsUpdate {
		log.Printf("File '%s' was updated, computing SHA256 hash and parsing lines...", filename)
		checksum.SetCachedFileChecksum(filename, computedFileChecksum)
		ParseLogFile(filename)
	}
}

func findLogFiles(root, ext string) []string {
	var filesList []string

	err := filepath.WalkDir(root, func(filename string, fileEntry fs.DirEntry, e error) error {
		if e != nil {
			log.Panicf("Error reading log file '%s' inside folder '%s': %s", filename, root, e)
		}

		if strings.ToLower(filepath.Ext(fileEntry.Name())) == strings.ToLower(ext) {
			filesList = append(filesList, filename)
		}

		return nil
	})

	if err != nil {
		log.Panicf("Error finding log files inside folder '%s': %s", root, err)
	}

	return filesList
}
