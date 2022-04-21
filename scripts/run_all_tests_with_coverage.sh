#!/bin/sh

. /go/src/github.com/ferdn4ndo/userver-logger-api/scripts/prepare_test_environment.sh

cd /go/src/github.com/ferdn4ndo/userver-logger-api/

echo "Running tests (with coverage)..."
go test -coverprofile="$DATA_FOLDER/$COVERAGE_DATA_FILE" ./... 

echo "Generating report coverage..."
go tool cover --func "$DATA_FOLDER/$COVERAGE_DATA_FILE" > "$DATA_FOLDER/$COVERAGE_REPORT_FILE"

echo "\nCOVERAGE REPORT (PER FUNCTION):\n"
cat "$DATA_FOLDER/$COVERAGE_REPORT_FILE"
echo ""

echo "Report exported to '$DATA_FOLDER/$COVERAGE_REPORT_FILE'!"
