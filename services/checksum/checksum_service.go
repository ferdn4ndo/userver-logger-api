package checksum

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ferdn4ndo/userver-logger-api/services/file"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

func ComputeFileChecksum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GetCachedChecksumFilename(fileFullPath string) string {
	tempFolder := file.GetTempFolder()

	filename := strings.TrimSuffix(filepath.Base(fileFullPath), filepath.Ext(fileFullPath))

	return fmt.Sprintf("%s/cached_checksum_%s.txt", tempFolder, filename)
}

func GetCachedFileChecksum(filename string) string {
	cachedFilename := GetCachedChecksumFilename(filename)
	if _, err := os.Stat(cachedFilename); errors.Is(err, os.ErrNotExist) {
		return ""
	}

	dat, err := os.ReadFile(cachedFilename)
	if err != nil {
		logging.Errorf("Error reading cached file %s", cachedFilename)
	}

	return string(dat)
}

func SetCachedFileChecksum(filename string, checksum string) {
	cachedFilename := GetCachedChecksumFilename(filename)
	fo, err := os.Create(cachedFilename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	fo.Write([]byte(checksum))
}
