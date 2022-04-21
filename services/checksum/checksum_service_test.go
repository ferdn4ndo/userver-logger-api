package checksum

import (
	"fmt"
	"testing"
)

func TestComputeFileChecksum(test *testing.T) {
	fileFullPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture/sample-app.log"
	expectedChecksum := "3c165740a1b151917f516e37a8f51cfdd72f01d27fbcd325d81cb75aa3dd47b8"

	computedChecksum, _ := ComputeFileChecksum(fileFullPath)
	if expectedChecksum != computedChecksum {
		test.Errorf(fmt.Sprintf("Failed asserting that the computed checksum '%s' is equal to expected '%s'.", computedChecksum, expectedChecksum))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}
