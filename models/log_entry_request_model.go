package models

import (
	"fmt"
	"net/http"
)

// LogEntryRequest is the request payload for LogEntry data model.
//
// NOTE: It's good practice to have well defined request and response payloads
// so you can manage the specific inputs and outputs for clients, and also gives
// you the opportunity to transform data on input or output, for example
// on request, we'd like to protect certain fields and on output perhaps
// we'd like to include a computed field based on other values that aren't
// in the data model. Also, check out this awesome blog post on struct composition:
// http://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
type LogEntryRequest struct {
	*LogEntry

	ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (logEntryRequest *LogEntryRequest) Bind(request *http.Request) error {
	if logEntryRequest.Producer == "" {
		return fmt.Errorf("The field 'producer' is required.")
	}

	if logEntryRequest.Message == "" {
		return fmt.Errorf("The field 'message' is required.")
	}

	// just a post-process after a decode..
	logEntryRequest.ProtectedID = "" // unset the protected ID

	return nil
}
