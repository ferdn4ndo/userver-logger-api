package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
)

type key string

const (
	PARAM_LOG_ENTRY_ID key = "logEntryId"
	CONTEXT_LOG_ENTRY  key = "contextLogEntry"
)

type LogEntrySingleController struct {
	DbService database.DatabaseServiceInterface
}

// Param converter (context) for the "logEntryId"
func (controller LogEntrySingleController) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logEntryId := chi.URLParam(request, "logEntryId")
		if logEntryId == "" {
			handler.RenderError(writer, request, handler.ServerErrorMsgRenderer("log entry ID is required"))

			return
		}

		id, err := strconv.Atoi(logEntryId)
		if err != nil {
			handler.RenderError(writer, request, handler.ServerErrorMsgRenderer("invalid log entry ID"))

			return
		}

		logEntryDbService := log_entry.LogEntryDatabaseService{DbService: controller.DbService}

		logEntryExists, err := logEntryDbService.CheckIfIdExists(uint(id))
		if err != nil {
			handler.RenderError(writer, request, handler.ServerErrorMsgRenderer("could not determine if log entry ID exists"))

			return
		}

		if !logEntryExists {
			handler.RenderError(writer, request, handler.ErrNotFound)

			return
		}

		logEntryPointer, err := logEntryDbService.GetLogEntryById(uint(id))
		if err != nil {
			handler.RenderError(writer, request, handler.ErrorRenderer(err))

			return
		}

		ctx := context.WithValue(request.Context(), CONTEXT_LOG_ENTRY, logEntryPointer)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

// GET /log-entries/{id}
func (controller LogEntrySingleController) Get(writer http.ResponseWriter, request *http.Request) {
	logEntryPointer := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)

	logEntryResponseService := log_entry.LogEntryResponseService{}
	error := render.Render(writer, request, logEntryResponseService.NewLogEntryResponse(logEntryPointer))
	if error != nil {
		handler.RenderError(writer, request, handler.ServerErrorRenderer(error))

		return
	}
}

// PUT /log-entries/{id}
func (controller LogEntrySingleController) Put(writer http.ResponseWriter, request *http.Request) {
	logEntryPointer := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)
	logEntryRequest := &models.LogEntryRequest{}

	if error := render.Bind(request, logEntryRequest); error != nil {
		handler.RenderError(writer, request, handler.ErrBadRequest)

		return
	}

	logEntryDbService := log_entry.LogEntryDatabaseService{DbService: controller.DbService}
	updatedLogEntryPointer, error := logEntryDbService.UpdateLogEntry((*logEntryPointer).ID, logEntryRequest)
	if error != nil {
		handler.RenderError(writer, request, handler.ServerErrorRenderer(error))

		return
	}

	logEntryResponseService := log_entry.LogEntryResponseService{}
	error = render.Render(writer, request, logEntryResponseService.NewLogEntryResponse(updatedLogEntryPointer))
	if error != nil {
		handler.RenderError(writer, request, handler.ServerErrorRenderer(error))

		return
	}
}

// DELETE /log-entries/{id}
func (controller LogEntrySingleController) Delete(writer http.ResponseWriter, request *http.Request) {
	logEntry := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)

	logEntryDbService := log_entry.LogEntryDatabaseService{DbService: controller.DbService}
	error := logEntryDbService.DeleteLogEntry(logEntry.ID)
	if error != nil {
		handler.RenderError(writer, request, handler.ServerErrorRenderer(error))

		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
