package http_request

import (
	"fmt"
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type MockedReadCloser struct{}

func (readerCloser MockedReadCloser) Reader() {
	fmt.Println("MockedReadCloser Reader")
}

func (readerCloser MockedReadCloser) Closer() {
	fmt.Println("MockedReadCloser Closer")
}

type MockedHttpResponseWriter struct {
	status int
}

func (httpRequest MockedHttpResponseWriter) Write(data []byte) (int, error) {
	fmt.Printf("%s", data)

	if httpRequest.status == 0 {
		httpRequest.status = 200
	}

	return httpRequest.status, nil
}

func (httpRequest MockedHttpResponseWriter) Header() http.Header {
	logging.Debug("Mock RealResponseWriter Header")

	var headers map[string][]string

	return headers
}

func (httpRequest MockedHttpResponseWriter) WriteHeader(statusCode int) {
	httpRequest.status = statusCode
}
