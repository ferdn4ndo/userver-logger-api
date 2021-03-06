package file

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type ConsumedLinesFileServiceInterface interface {
	GetLastConsumedLinesFilePath()
	UpdateLastConsumedLines() error
}

type ConsumedLinesFileService struct {
	LogFilePath string
}

func (service ConsumedLinesFileService) GetProducer() string {
	if service.LogFilePath == "" {
		logging.Errorf("Unable to start the service without a log file path (LogFilePath is empty).")

		return "<unknown>"
	}

	return GetContainerNameFromPath(service.LogFilePath)
}

func (service ConsumedLinesFileService) GetLastConsumedLinesFilePath() string {
	producer := service.GetProducer()

	folder := fmt.Sprintf("%s/last-consumed", GetTempFolder())
	createFolderIfNotExists(folder)

	filepath := fmt.Sprintf("%s/%s.last-consumed.log", folder, producer)

	_, err := os.Stat(filepath)
	if errors.Is(err, os.ErrNotExist) {
		cpCommand := exec.Command("touch", filepath)
		if err := cpCommand.Run(); err != nil {
			logging.Errorf("Error creating empty last consumed file at '%s': %s", filepath, err)
		}
	} else if err != nil {
		logging.Errorf("Error checking empty last consumed file at '%s': %s", filepath, err)
	}

	return filepath
}

func (service ConsumedLinesFileService) UpdateLastConsumedLines() error {
	destFolder := service.GetLastConsumedLinesFilePath()
	cpCommand := exec.Command("cp", "-f", service.LogFilePath, destFolder)
	err := cpCommand.Run()

	return err
}

func (service ConsumedLinesFileService) GetLastConsumedFileDiff() (string, error) {
	consumedFilePath := service.GetLastConsumedLinesFilePath()
	logging.Debugf("Reading consumed log path '%s'", consumedFilePath)
	if _, err := os.Stat(consumedFilePath); err != nil {
		return "", err
	}

	logging.Debugf("Comparing files '%s' and '%s'...", service.LogFilePath, consumedFilePath)
	cmd := exec.Command("comm", "--nocheck-order", "-3", service.LogFilePath, consumedFilePath)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

func (service ConsumedLinesFileService) IsLastConsumedFileDifferent() (bool, error) {
	diff, err := service.GetLastConsumedFileDiff()
	if err != nil {
		return false, err
	}

	return string(diff) != "", nil
}

func (service ConsumedLinesFileService) RequiresUpdate() bool {
	isDifferent, err := service.IsLastConsumedFileDifferent()

	if errors.Is(err, os.ErrNotExist) {
		return true
	} else if err != nil {
		logging.Errorf("Error when comparing consumed file: %s", err)
	}

	return isDifferent
}
