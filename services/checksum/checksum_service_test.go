package checksum

import (
	"fmt"
	"testing"
)

func TestComputeFileChecksum(test *testing.T) {
	fileFullPath := "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture/sample-app.log"
	expectedChecksum := "4c2a6ae0d9ac1c01e8f8f0d418830b6486ed4bab806f30d829e2af71348cc61d"

	computedChecksum, _ := ComputeFileChecksum(fileFullPath)
	if expectedChecksum != computedChecksum {
		test.Errorf(fmt.Sprintf("Failed asserting that the computed checksum '%s' is equal to expected '%s'.", computedChecksum, expectedChecksum))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}
