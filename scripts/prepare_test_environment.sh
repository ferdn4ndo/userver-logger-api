#!/bin/sh

export VIRTUAL_HOST=userver-logger-api.test.lan
export BASIC_AUTH_USERNAME=test
export BASIC_AUTH_PASSWORD=@StR0ng!P4ssW0rD+
export SERVER_PORT=8888
export LOG_FILES_FOLDER=/log_files
export TMP_FOLDER=/go/src/github.com/ferdn4ndo/userver-logger-api/tmp
export DATA_FOLDER=/go/src/github.com/ferdn4ndo/userver-logger-api/data
export FIXTURE_FOLDER=/go/src/github.com/ferdn4ndo/userver-logger-api/fixture
export DATABASE_FILE=sqlite.db
export TEST_DATABASE_FILE=test.sqlite.db
export EMPTY_DATABASE_FILE=empty.sqlite.db

export COVERAGE_DATA_FILE=coverage.out
export COVERAGE_REPORT_FILE=coverage.report.txt

/go/src/github.com/ferdn4ndo/userver-logger-api/scripts/reset_test_database.sh

cd /go/src/github.com/ferdn4ndo/userver-logger-api/

go get -u -t ./...

go mod download

go mod tidy
