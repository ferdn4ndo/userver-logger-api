package authentication

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
)

func BasicAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Extract the username and password from the request
		// Authorization header. If no Authentication header is present
		// or the header value is invalid, then the 'ok' return value
		// will be false.
		username, password, ok := request.BasicAuth()
		if ok {
			credentialsAreValid := validateCredentials(username, password)

			// If the username and password are correct, then call
			// the next handler in the chain. Make sure to return
			// afterwards, so that none of the code below is run.
			if credentialsAreValid {
				next.ServeHTTP(writer, request)
				return
			}
		}

		// If the Authentication header is not present, is invalid, or the
		// username or password is wrong, then set a WWW-Authenticate
		// header to inform the client that we expect them to use basic
		// authentication and send a 401 Unauthorized response.
		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

		handler.UnauthorizedHandler(writer)
	})
}

func validateCredentials(username string, password string) bool {
	// Calculate SHA-256 hashes for the provided and expected
	// usernames and passwords.
	usernameHash := computeSha256(username)
	passwordHash := computeSha256(password)
	expectedUsernameHash := computeSha256(getAuthUsername())
	expectedPasswordHash := computeSha256(getAuthPassword())

	// Use the subtle.ConstantTimeCompare() function to check if
	// the provided username and password hashes equal the
	// expected username and password hashes. ConstantTimeCompare
	// will return 1 if the values are equal, or 0 otherwise.
	// Importantly, we should to do the work to evaluate both the
	// username and password before checking the return values to
	// avoid leaking information.
	usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
	passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

	return usernameMatch && passwordMatch
}

func getAuthUsername() string {
	username := environment.GetEnvKey("BASIC_AUTH_USERNAME")

	if username == "" {
		log.Fatal("Unable to determine the basic authentication username!")
	}

	return username
}

func getAuthPassword() string {
	password := environment.GetEnvKey("BASIC_AUTH_PASSWORD")

	if password == "" {
		log.Fatal("Unable to determine the basic authentication password!")
	}

	return password
}

func computeSha256(input string) [32]byte {
	return sha256.Sum256([]byte(input))
}
