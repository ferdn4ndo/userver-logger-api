package pagination

import (
	"testing"
)

func TestComputePreviousPageOffset(test *testing.T) {
	if computePreviousPageOffset(300, 100) != 200 {
		test.Fatal("Failed asserting that the previous page offset will be 200 with size = 100 and current offset = 200.")
	}

	if computePreviousPageOffset(50, 100) != 0 {
		test.Fatal("Failed asserting that the previous page offset will be 0 with size = 100 and current offset = 50.")
	}

	test.Log("Finished testing the computePreviousPageOffset() method")
}
