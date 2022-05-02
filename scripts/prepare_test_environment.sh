#!/bin/sh

export VIRTUAL_HOST=userver-logger-api.test.lan
export BASIC_AUTH_USERNAME=test
export BASIC_AUTH_PASSWORD=@StR0ng!P4ssW0rD+
export TEST_DATABASE_FILE=test.sqlite.db

export COVERAGE_DATA_FILE=coverage.out
export COVERAGE_REPORT_FILE=coverage.xml

/go/src/github.com/ferdn4ndo/userver-logger-api/scripts/reset_test_database.sh

cd /go/src/github.com/ferdn4ndo/userver-logger-api/

go get -u -t ./...

go mod download

go mod tidy
