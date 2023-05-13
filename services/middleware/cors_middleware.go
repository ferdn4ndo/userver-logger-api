package middleware

import (
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

const corsHeader = "Access-Control-Allow-Origin"

func AddCorsHeader(next http.Handler) http.Handler {
	corsValue := getCorsHeaderValue(environment.GetEnvKey("CORS_ALLOWED_HOSTS"))

	handler := func(writer http.ResponseWriter, request *http.Request) {
		context := request.Context()
		if writer.Header().Get(corsHeader) == "" {
			writer.Header().Add(
				corsHeader,
				corsValue,
			)
		}
		next.ServeHTTP(writer, request.WithContext(context))
	}

	return http.HandlerFunc(handler)
}

func getCorsHeaderValue(corsValue string) string {
	if corsValue == "" {
		logging.Warning("The CORS_ALLOWED_HOSTS env var is empty, meaning that the API will be accessible from any host (assuming the '*' wildcard)!")

		return "*"
	}

	return corsValue
}
