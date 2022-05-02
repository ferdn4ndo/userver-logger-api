package models

import "net/http"

type LogEntryList struct {
	LogEntries []*LogEntry `json:"log_entries" gorm:"-"`
}

func (*LogEntryList) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
