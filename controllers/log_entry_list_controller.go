package controllers

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
	"github.com/ferdn4ndo/userver-logger-api/services/pagination"
)

type LogEntryListController struct {
	DbService database.DatabaseServiceInterface
}

// GET /log-entries
func (controller LogEntryListController) Get(writer http.ResponseWriter, request *http.Request) {
	paginationService := pagination.PaginationService{Request: request}
	offset, limit := paginationService.GetRequestOffsetAndLimit()
	searchParams := log_entry.GetLogEntrySearchParams(request)
	logEntryDbService := log_entry.LogEntryDatabaseService{DbService: controller.DbService}

	logEntryListPointer, totalResults, error := logEntryDbService.GetAllLogEntries(request, searchParams)
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}

	logEntryResponseService := log_entry.LogEntryResponseService{PaginationService: paginationService}
	error = render.Render(writer, request, logEntryResponseService.NewLogEntryListResponse(logEntryListPointer, offset, limit, totalResults))
	if error != nil {
		render.Render(writer, request, handler.ErrorRenderer(error))
	}
}

// POST /log-entries
func (controller LogEntryListController) Post(writer http.ResponseWriter, request *http.Request) {
	logEntryRequest := &models.LogEntryRequest{}

	if error := render.Bind(request, logEntryRequest); error != nil {
		render.Render(writer, request, handler.ErrBadRequest)

		return
	}

	logEntryDbService := log_entry.LogEntryDatabaseService{DbService: controller.DbService}
	if error := logEntryDbService.AddLogEntry(logEntryRequest.LogEntry); error != nil {
		render.Render(writer, request, handler.ErrorRenderer(error))

		return
	}

	logEntryResponseService := log_entry.LogEntryResponseService{}
	render.Status(request, http.StatusCreated)
	error := render.Render(writer, request, logEntryResponseService.NewLogEntryResponse(logEntryRequest.LogEntry))
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}
}
