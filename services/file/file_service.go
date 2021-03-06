package file

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

func createFolderIfNotExists(folderPath string) {
	// If path is already a directory, MkdirAll does nothing and returns nil.
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		logging.Errorf("Error while creating folder at '%s': %s", folderPath, err)
	}
}

func GetTempFolder() string {
	tempFolder := environment.GetEnvKey("TMP_FOLDER")

	createFolderIfNotExists(tempFolder)

	return tempFolder
}

func GetDataFolder() string {
	dataFolder := environment.GetEnvKey("DATA_FOLDER")

	createFolderIfNotExists(dataFolder)

	return dataFolder

}

func GetFixtureFolder() string {
	fixtureFolder := environment.GetEnvKey("FIXTURE_FOLDER")

	createFolderIfNotExists(fixtureFolder)

	return fixtureFolder

}

func GetLogFilesFolder() string {
	logFilesFolder := environment.GetEnvKey("LOG_FILES_FOLDER")

	createFolderIfNotExists(logFilesFolder)

	return logFilesFolder
}

func GetContainerNameFromPath(fullFilePath string) string {
	filename := filepath.Base(fullFilePath)
	extension := filepath.Ext(filename)

	return strings.TrimSuffix(filename, extension)
}
