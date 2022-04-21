#!/bin/sh

. /go/src/github.com/ferdn4ndo/userver-logger-api/scripts/prepare_test_environment.sh

cd /go/src/github.com/ferdn4ndo/userver-logger-api/

go test -coverprofile="$DATA_FOLDER/$COVERAGE_DATA_FILE" ./... > /dev/null

coverage=$(go tool cover -func "$DATA_FOLDER/$COVERAGE_DATA_FILE" | grep total | awk '{print substr($3, 1, length($3)-1)}')
echo "$coverage"
