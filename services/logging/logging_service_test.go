package logging

import (
	"strconv"
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

func TestGetCurrentLogLevel(test *testing.T) {
	expectedLevelText := environment.GetEnvKey("INTERNAL_LOG_LEVEL")
	expectedLevel, _ := strconv.Atoi(expectedLevelText)

	computedLevel := getCurrentLogLevel()

	if computedLevel != expectedLevel {
		test.Fatalf("Failed asserting that the computed log level '%d' matches the expected '%d'", computedLevel, expectedLevel)
	}
}

func TestPrintMessageF(test *testing.T) {
	expectedOutput := "HELLO WORLD!"

	realOutput := printMessageF(expectedOutput)

	if string(realOutput) != expectedOutput {
		test.Fatalf("Failed asserting that the printed message '%s' matches the expected '%s'", realOutput, expectedOutput)
	}
}

func TestPrintMessageIfLevelAllows(test *testing.T) {
	currentLevel := getCurrentLogLevel()

	newLevel := currentLevel - 1
	output := printMessageIfLevelAllows(newLevel, "test")
	if !output {
		test.Fatalf("Failed asserting that the message would be written using level %d (current level: %d)", newLevel, currentLevel)
	}

	newLevel = currentLevel + 1
	output = printMessageIfLevelAllows(newLevel, "test")
	if output {
		test.Fatalf("Failed asserting that the message would NOT be written using level %d (current level: %d)", newLevel, currentLevel)
	}
}
