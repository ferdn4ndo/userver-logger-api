package checksum

import (
	"fmt"
	"testing"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

func TestComputeFileChecksum(test *testing.T) {
	fileFullPath := fmt.Sprintf("%s/sample-app.log", environment.GetEnvKey("FIXTURE_FOLDER"))
	expectedChecksum := "4c2a6ae0d9ac1c01e8f8f0d418830b6486ed4bab806f30d829e2af71348cc61d"

	computedChecksum, _ := ComputeFileChecksum(fileFullPath)
	if expectedChecksum != computedChecksum {
		test.Fatalf(fmt.Sprintf("Failed asserting that the computed checksum '%s' is equal to expected '%s'.", computedChecksum, expectedChecksum))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}
