package log_entry

import (
	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/pagination"
)

type LogEntryResponseServiceInterface interface {
	NewLogEntryResponse(logEntry *models.LogEntry) *models.LogEntryResponse
	NewLogEntryListResponse(logEntries *models.LogEntryList, offset int, limit int, totalCount int) *pagination.PaginatedResponse
}

type LogEntryResponseService struct {
	PaginationService pagination.PaginationServiceInterface
}

func (service LogEntryResponseService) NewLogEntryResponse(logEntry *models.LogEntry) *models.LogEntryResponse {
	resp := &models.LogEntryResponse{LogEntry: logEntry}

	return resp
}

func (service LogEntryResponseService) NewLogEntryListResponse(logEntries *models.LogEntryList, offset int, limit int, totalCount int) *pagination.PaginatedResponse {
	list := []render.Renderer{}
	for _, logEntry := range logEntries.LogEntries {
		list = append(list, service.NewLogEntryResponse(logEntry))
	}

	paginationResponse := service.PaginationService.PreparePaginatedResponse(list, totalCount)

	return &paginationResponse
}
