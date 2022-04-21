#!/bin/sh

echo "Check if any log is present"
LOG_FILES_COUNT=$(ls "$LOG_FILES_FOLDER" | wc -l)

if [ "$LOG_FILES_COUNT" -eq "0" ]; then
    echo "No log file is present, importing the fixture one"
    cp /go/src/github.com/ferdn4ndo/userver-logger-api/fixture/sample-app.log "$LOG_FILES_FOLDER/sample-app.log"
else
    echo "$LOG_FILES_COUNT log file(s) found. Skipping fixture import."
fi

echo "Starting AIR"
air
