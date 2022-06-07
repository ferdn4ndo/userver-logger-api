package logging

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

const LOG_LEVEL_NONE = 0
const LOG_LEVEL_ERROR = 25
const LOG_LEVEL_WARNING = 50
const LOG_LEVEL_INFO = 75
const LOG_LEVEL_DEBUG = 100

// Retrieve the current log level based on the 'INTERNAL_LOG_LEVEL' env var
func getCurrentLogLevel() int {
	level_text := environment.GetEnvKey("INTERNAL_LOG_LEVEL")

	level, err := strconv.Atoi(level_text)
	if err != nil {
		log.Fatalf("Error getting log level: '%s'", err)
	}

	return level
}

// Logs an error message (in case the current log level includes it)
func Error(message string) {
	if getCurrentLogLevel() >= LOG_LEVEL_ERROR {
		log.Printf("[ERROR] %s\n", message)
	}
}

// A Sprintf-like method for an Error level message
func Errorf(message string, arguments ...any) {
	Error(fmt.Sprintf(message, arguments...))
}

// Logs a warning message (in case the current log level includes it)
func Warning(message string) {
	if getCurrentLogLevel() >= LOG_LEVEL_WARNING {
		log.Printf("[WARNING] %s\n", message)
	}
}

// A Sprintf-like method for a Warning level message
func Warningf(message string, arguments ...any) {
	Warning(fmt.Sprintf(message, arguments...))
}

// Logs an info message (in case the current log level includes it)
func Info(message string) {
	if getCurrentLogLevel() >= LOG_LEVEL_INFO {
		log.Printf("[INFO] %s\n", message)
	}
}

// A Sprintf-like method for an Info level message
func Infof(message string, arguments ...any) {
	Info(fmt.Sprintf(message, arguments...))
}

// Logs a debug message (in case the current log level includes it)
func Debug(message string) {
	if getCurrentLogLevel() >= LOG_LEVEL_DEBUG {
		log.Printf("[DEBUG] %s\n", message)
	}
}

// A Sprintf-like method for a Debug level message
func Debugf(message string, arguments ...any) {
	Debug(fmt.Sprintf(message, arguments...))
}
