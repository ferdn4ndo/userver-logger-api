package logging

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

const (
	LOG_LEVEL_NONE    = 0
	LOG_LEVEL_ERROR   = 25
	LOG_LEVEL_WARNING = 50
	LOG_LEVEL_INFO    = 75
	LOG_LEVEL_DEBUG   = 100
)

// Retrieve the current log level based on the 'INTERNAL_LOG_LEVEL' env var
func getCurrentLogLevel() int {
	level_text := environment.GetEnvKey("INTERNAL_LOG_LEVEL")

	level, err := strconv.Atoi(level_text)
	if err != nil {
		log.Fatalf("Error getting log level: '%s'", err)
	}

	return level
}

// Prints a given message in printf-like (can receive multiple arguments)
func printMessageF(format string, arguments ...any) string {
	message := fmt.Sprintf(format, arguments...)

	log.Print(message)

	return message
}

// Prints a given message (printf-like) if the current log level is bigger or equal to min_level
func printMessageIfLevelAllows(min_level int, message string, arguments ...any) bool {
	if getCurrentLogLevel() >= min_level {
		printMessageF(message, arguments...)

		return true
	}

	return false
}

// Logs an error message (in case the current log level includes it)
func Error(message string) bool {
	return printMessageIfLevelAllows(LOG_LEVEL_ERROR, "[ERROR] %s\n", message)
}

// A Sprintf-like method for an Error level message
func Errorf(message string, arguments ...any) bool {
	return Error(fmt.Sprintf(message, arguments...))
}

// Logs a warning message (in case the current log level includes it)
func Warning(message string) bool {
	return printMessageIfLevelAllows(LOG_LEVEL_WARNING, "[WARNING] %s\n", message)
}

// A Sprintf-like method for a Warning level message
func Warningf(message string, arguments ...any) bool {
	return Warning(fmt.Sprintf(message, arguments...))
}

// Logs an info message (in case the current log level includes it)
func Info(message string) bool {
	return printMessageIfLevelAllows(LOG_LEVEL_INFO, "[INFO] %s\n", message)
}

// A Sprintf-like method for an Info level message
func Infof(message string, arguments ...any) bool {
	return Info(fmt.Sprintf(message, arguments...))
}

// Logs a debug message (in case the current log level includes it)
func Debug(message string) bool {
	return printMessageIfLevelAllows(LOG_LEVEL_DEBUG, "[DEBUG] %s\n", message)
}

// A Sprintf-like method for a Debug level message
func Debugf(message string, arguments ...any) bool {
	return Debug(fmt.Sprintf(message, arguments...))
}
