package models

import "net/http"

type LogEntryList struct {
	LogEntries []*LogEntry `json:"log_entries" gorm:"-"`
}

func (*LogEntryList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
