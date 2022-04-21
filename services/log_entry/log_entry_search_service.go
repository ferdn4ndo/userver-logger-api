package log_entry

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type LogEntrySearchParams struct {
	Producer string
	Message  string
}

func GetLogEntrySearchParams(request *http.Request) *LogEntrySearchParams {
	searchParams := &LogEntrySearchParams{
		Producer: request.URL.Query().Get("producer"),
		Message:  request.URL.Query().Get("message"),
	}

	return searchParams
}

func ApplyLogEntryQuerySearchParams(query *gorm.DB, params *LogEntrySearchParams) *gorm.DB {

	if params.Producer != "" {
		query.Where("producer = ?", params.Producer)
	}

	if params.Message != "" {
		query.Where("message LIKE ?", fmt.Sprintf("%%%s%%", params.Message))
	}

	return query
}
