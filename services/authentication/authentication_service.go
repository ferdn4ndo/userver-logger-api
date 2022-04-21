package authentication

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
)

func BasicAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the username and password from the request
		// Authorization header. If no Authentication header is present
		// or the header value is invalid, then the 'ok' return value
		// will be false.
		username, password, ok := r.BasicAuth()
		if ok {
			credentialsAreValid := validateCredentials(username, password)

			// If the username and password are correct, then call
			// the next handler in the chain. Make sure to return
			// afterwards, so that none of the code below is run.
			if credentialsAreValid {
				next.ServeHTTP(w, r)
				return
			}
		}

		// If the Authentication header is not present, is invalid, or the
		// username or password is wrong, then set a WWW-Authenticate
		// header to inform the client that we expect them to use basic
		// authentication and send a 401 Unauthorized response.
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

		handler.UnauthorizedHandler(w)
	})
}

func validateCredentials(username string, password string) bool {
	// Calculate SHA-256 hashes for the provided and expected
	// usernames and passwords.
	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))
	expectedUsernameHash := sha256.Sum256([]byte(environment.GetEnvKey("BASIC_AUTH_USERNAME")))
	expectedPasswordHash := sha256.Sum256([]byte(environment.GetEnvKey("BASIC_AUTH_PASSWORD")))

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
