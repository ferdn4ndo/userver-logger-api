package models

import "net/http"

type LogEntryResponse struct {
	*LogEntry

	// We could add an additional field to the response here, such as this
	// elapsed computed property:
	//Elapsed int64 `json:"elapsed"`
}

func (logEntry *LogEntryResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	//logEntry.Elapsed = 10
	return nil
}
