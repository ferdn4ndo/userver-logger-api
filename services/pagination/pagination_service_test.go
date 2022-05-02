package pagination

import (
	"testing"
)

func TestComputePreviousPageOffset(test *testing.T) {
	if 200 != computePreviousPageOffset(300, 100) {
		test.Fatal("Failed asserting that the previous page offset will be 200 with size = 100 and current offset = 200.")
	}

	if 0 != computePreviousPageOffset(50, 100) {
		test.Fatal("Failed asserting that the previous page offset will be 0 with size = 100 and current offset = 50.")
	}

	test.Log("Finished testing the computePreviousPageOffset() method")
}
