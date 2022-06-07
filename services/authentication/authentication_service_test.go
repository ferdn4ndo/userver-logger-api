package authentication

import (
	"testing"
)

func TestValidateCredentials(test *testing.T) {
	validUsername := "test"
	validPassword := "@StR0ng!P4ssW0rD+"

	validResult := validateCredentials(validUsername, validPassword)
	if validResult == false {
		test.Fatalf("Failed asserting that the valid credentials are authorized successfully!")
	}

	invalidResult := validateCredentials("rand0m", "cr3d3nt14ls")
	if invalidResult == true {
		test.Fatalf("Failed asserting that the invalid credentials are not authorized successfully!")
	}

	test.Log("Finished testing the ValidateCredentials() method")
}
