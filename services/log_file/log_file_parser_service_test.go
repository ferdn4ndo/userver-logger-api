package log_file

import (
	"fmt"
	"testing"
)

func TestGetContainerNameFromPath(test *testing.T) {
	filePath := "/test/path/container-name.log"
	expectedName := "container-name"
	computedName := getContainerNameFromPath(filePath)

	if expectedName != computedName {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed container name '%s' is equal to the expected '%s'.", computedName, expectedName))
	}

	test.Log("Finished testing the getContainerNameFromPath() method")
}
