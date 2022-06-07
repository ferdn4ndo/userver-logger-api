#!/bin/sh

. /go/src/github.com/ferdn4ndo/userver-logger-api/scripts/prepare_test_environment.sh

cd /go/src/github.com/ferdn4ndo/userver-logger-api/

echo "Running tests..."
go test -race ./...
