package application

import (
	"fmt"
	"testing"
)

func TestGetServerPort(test *testing.T) {
	expectedPort := 5000

	appObject := Application{}
	appPort := appObject.getServerPort()

	if expectedPort != appPort {
		test.Fatalf(fmt.Sprintf("Failed asserting that the server port '%d' matches the expected '%d'.", appPort, expectedPort))
	}

	test.Log("Finished testing the GetBaseApplication() method")
}
