package log_entry

import (
	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/pagination"
)

func NewLogEntryResponse(logEntry *models.LogEntry) *models.LogEntryResponse {
	resp := &models.LogEntryResponse{LogEntry: logEntry}

	return resp
}

func NewLogEntryListResponse(logEntries *models.LogEntryList, offset int, limit int, totalCount int) *pagination.PaginatedResponse {
	list := []render.Renderer{}
	for _, logEntry := range logEntries.LogEntries {
		list = append(list, NewLogEntryResponse(logEntry))
	}

	paginationResponse := pagination.PreparePaginatedResponse(list, offset, limit, totalCount)

	return &paginationResponse
}
