package environment

import (
	"fmt"
	"testing"
)

func checkEnvVarData(test *testing.T, envVar EnvVar, expectedKey string, expectedDescription string, expectedRequired bool, expectedDefaultValue string) {
	if envVar.Key != expectedKey {
		test.Errorf(fmt.Sprintf("Failed asserting that the attribute 'Key' has the value '%s' (got '%s' instead).", expectedKey, envVar.Key))
	}

	if envVar.Description != expectedDescription {
		test.Errorf(fmt.Sprintf("Failed asserting that the attribute 'Description' has the value '%s' (got '%s' instead).", expectedDescription, envVar.Description))
	}

	if envVar.Required != expectedRequired {
		test.Errorf(fmt.Sprintf("Failed asserting that the attribute 'Required' has the value '%t' (got '%t' instead).", expectedRequired, envVar.Required))
	}

	if envVar.DefaultValue != expectedDefaultValue {
		test.Errorf(fmt.Sprintf("Failed asserting that the attribute 'DefaultValue' has the value '%s' (got '%s' instead).", expectedDefaultValue, envVar.DefaultValue))
	}
}

func TestEnvVarBasicAttributes(test *testing.T) {
	expectedKey := "foo"
	expectedDescription := "bar"
	expectedRequired := true
	expectedDefaultValue := "default"
	expectedCurrentValue := "current"

	envVar := EnvVar{
		Key:          expectedKey,
		Description:  expectedDescription,
		Required:     expectedRequired,
		DefaultValue: expectedDefaultValue,
		CurrentValue: expectedCurrentValue,
	}

	checkEnvVarData(test, envVar, expectedKey, expectedDescription, expectedRequired, expectedDefaultValue)

	if envVar.CurrentValue != expectedCurrentValue {
		test.Errorf(fmt.Sprintf("Failed asserting that the attribute 'CurrentValue' has the value '%s' (got '%s' instead).", expectedCurrentValue, envVar.CurrentValue))
	}

	test.Log("Finished testing the ComputeFileChecksum() method")
}

func TestAddVariableToList(test *testing.T) {
	expectedKey := "foo"
	expectedDescription := "bar"
	expectedRequired := true
	expectedDefaultValue := "default"
	list := EnvVarList{}

	list.addVariable(expectedKey, expectedDescription, expectedRequired, expectedDefaultValue)

	totalVariables := len(list.Variables)
	if totalVariables != 1 {
		test.Errorf(fmt.Sprintf("Failed asserting that one varialbe was added to the list! (count: %d)", totalVariables))
	}

	checkEnvVarData(test, list.Variables[0], expectedKey, expectedDescription, expectedRequired, expectedDefaultValue)
}

func TestEnvHasKey(test *testing.T) {
	expectedKey := "foo"
	expectedDescription := "bar"
	expectedRequired := true
	expectedDefaultValue := "default"
	list := EnvVarList{}

	list.addVariable(expectedKey, expectedDescription, expectedRequired, expectedDefaultValue)

	hasKeyThatExists := list.envHasKey(expectedKey)
	if !hasKeyThatExists {
		test.Fatalf("Failed asserting that envHasKey method succeeded with an existing key.")
	}
	hasKeyThatDoesNotExist := list.envHasKey("random_key")
	if hasKeyThatDoesNotExist {
		test.Fatalf("Failed asserting that envHasKey method failed with a non-existing key.")
	}

	test.Log("Finished testing the envHasKey() method")
}
