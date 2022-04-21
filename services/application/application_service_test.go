package application

import (
	"fmt"
	"testing"
)

func TestGetBaseApplication(test *testing.T) {
	expectedUsername := "test"
	expectedPassword := "@StR0ng!P4ssW0rD+"
	expectedPort := 8888

	appObject := GetBaseApplication()

	if expectedUsername != appObject.Auth.Username {
		test.Errorf(fmt.Sprintf("Failed asserting that the auth username '%s' matches the expected '%s'.", appObject.Auth.Username, expectedUsername))
	}

	if expectedPassword != appObject.Auth.Password {
		test.Errorf(fmt.Sprintf("Failed asserting that the auth username '%s' matches the expected '%s'.", appObject.Auth.Username, expectedUsername))
	}

	if expectedPort != appObject.Port {
		test.Errorf(fmt.Sprintf("Failed asserting that the auth username '%s' matches the expected '%s'.", appObject.Auth.Username, expectedUsername))
	}

	test.Log("Finished testing the GetBaseApplication() method")
}
