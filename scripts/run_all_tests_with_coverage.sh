#!/bin/sh

. /go/src/github.com/ferdn4ndo/userver-logger-api/scripts/prepare_test_environment.sh

cd /go/src/github.com/ferdn4ndo/userver-logger-api/

echo "Installing dependencies for coverage reporting..."
go install github.com/axw/gocov/gocov@latest
go install github.com/AlekSi/gocov-xml@latest

echo "Running tests (with coverage)..."
go test -race -covermode=atomic -coverprofile="$DATA_FOLDER/$COVERAGE_DATA_FILE" ./... 

echo "Raw coverage output exported to '$DATA_FOLDER/$COVERAGE_DATA_FILE'!"
