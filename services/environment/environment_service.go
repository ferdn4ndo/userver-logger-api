package environment

import (
	"log"
	"os"
)

type EnvVar struct {
	Key          string
	Description  string
	Required     bool
	DefaultValue string
	CurrentValue string
}

type EnvVarList struct {
	Variables []EnvVar
}

func (environment *EnvVarList) addVariable(key string, description string, required bool, defaultValue string) {
	if environment.envHasKey(key) {
		log.Panicf("Environment key '%s' was already set!\n", key)
	}

	variable := EnvVar{
		Key:          key,
		Description:  description,
		Required:     required,
		DefaultValue: defaultValue,
	}

	environment.Variables = append(environment.Variables, variable)
}

func (environment *EnvVarList) envHasKey(key string) bool {
	for _, variable := range environment.Variables {
		if variable.Key == key {
			return true
		}
	}

	return false
}

func (environment *EnvVarList) getVariableByKey(key string) *EnvVar {
	if !environment.envHasKey(key) {
		log.Panicf("Trying to fetch unknown environment variable '%s'!\n", key)
	}

	for _, variable := range environment.Variables {
		if variable.Key == key {
			return &variable
		}
	}

	return nil
}

func (environment *EnvVarList) fetchKeyValue(key string) string {
	variable := *environment.getVariableByKey(key)

	envValue := os.Getenv(key)

	if envValue == "" && variable.Required && variable.DefaultValue == "" {
		log.Panicf("The required environment variable '%s' is not set!\n", key)
	}

	if envValue != "" {
		variable.CurrentValue = envValue
	} else if variable.DefaultValue != "" {
		variable.CurrentValue = variable.DefaultValue
	}

	return variable.CurrentValue
}

func PreloadVariables() *EnvVarList {
	variables := EnvVarList{}

	variables.addVariable("VIRTUAL_HOST", "The virtual hostname to use if you're running the container behind a reverse proxy.", false, "localhost")
	variables.addVariable("LETSENCRYPT_HOST", "The virtual hostname to use in the SSL certificate generation by Let's Encrypt if you're running the container behind a reverse proxy.", false, "localhost")
	variables.addVariable("LETSENCRYPT_EMAIL", "The hostmaster e-mail to use in the SSL certificate generation by Let's Encrypt if you're running the container behind a reverse proxy.", false, "localhost")
	variables.addVariable("BASIC_AUTH_USERNAME", "The username to use in the Basic Authentication of the API endpoints.", true, "")
	variables.addVariable("BASIC_AUTH_PASSWORD", "The password to use in the Basic Authentication of the API endpoints.", true, "")
	variables.addVariable("INTERNAL_LOG_LEVEL", "The minimum log level to be printed (for internal api workflows, not for the monitored containers).", false, "50")
	variables.addVariable("LOG_FILES_FOLDER", "The location of the log files to be watched.", false, "/log_files")
	variables.addVariable("TMP_FOLDER", "The location of the temporary files used while running the service.", false, "/go/src/github.com/ferdn4ndo/userver-logger-api/tmp")
	variables.addVariable("DATA_FOLDER", "The location of the data files used while running the service.", false, "/go/src/github.com/ferdn4ndo/userver-logger-api/data")
	variables.addVariable("FIXTURE_FOLDER", "The location of the fixture files for preloading internal service data", false, "/go/src/github.com/ferdn4ndo/userver-logger-api/fixture")
	variables.addVariable("DATABASE_FILE", "The filename of the SQLite database file (inside the data folder) to store the parsed log entries.", false, "sqlite.db")
	variables.addVariable("TEST_DATABASE_FILE", "The filename of the SQLite database file (inside the data folder) to use during the tests.", false, "test.sqlite.db")
	variables.addVariable("EMPTY_DATABASE_FILE", "The filename of the SQLite database file (inside the fixture folder) without any table, to be used when preparing a new test environment.", false, "empty.sqlite.db")
	variables.addVariable("FILE_SCAN_INTERVAL", "The interval (in seconds) to check for changes in the log files (for parsing).", false, "5")

	return &variables
}

func GetEnvKey(key string) string {
	variables := PreloadVariables()

	return variables.fetchKeyValue(key)
}
