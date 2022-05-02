package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/models"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
	"github.com/ferdn4ndo/userver-logger-api/services/pagination"
)

var PARAM_LOG_ENTRY_ID = "logEntryId"
var CONTEXT_LOG_ENTRY = "contextLogEntry"

// Param converter (context) for the "logEntryId"
func LogEntryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logEntryId := chi.URLParam(request, "logEntryId")
		if logEntryId == "" {
			render.Render(writer, request, handler.ErrorRenderer(fmt.Errorf("Log entry ID is required.")))

			return
		}

		id, error := strconv.Atoi(logEntryId)
		if error != nil {
			render.Render(writer, request, handler.ErrorRenderer(fmt.Errorf("Invalid log entry ID.")))

			return
		}

		logEntryPointer, error := log_entry.GetLogEntryById(uint(id))
		if error != nil {
			if error == database.ErrNoMatch {
				render.Render(writer, request, handler.ErrNotFound)
			} else {
				render.Render(writer, request, handler.ErrorRenderer(error))
			}

			return
		}

		log.Printf("VALUE: %+v \n", logEntryPointer)
		log.Printf("VALUE2: %+v \n", *logEntryPointer)

		ctx := context.WithValue(request.Context(), CONTEXT_LOG_ENTRY, logEntryPointer)

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

// GET /log-entries
func GetListLogEntry(writer http.ResponseWriter, request *http.Request) {
	offset, limit := pagination.GetRequestOffsetAndLimit(request)
	searchParams := log_entry.GetLogEntrySearchParams(request)

	logEntryListPointer, totalResults, error := log_entry.GetAllLogEntries(limit, offset, searchParams)
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}

	error = render.Render(writer, request, log_entry.NewLogEntryListResponse(logEntryListPointer, offset, limit, totalResults))
	if error != nil {
		render.Render(writer, request, handler.ErrorRenderer(error))
	}
}

// POST /log-entries
func PostLogEntry(writer http.ResponseWriter, request *http.Request) {
	logEntryRequest := &models.LogEntryRequest{}

	if error := render.Bind(request, logEntryRequest); error != nil {
		render.Render(writer, request, handler.ErrBadRequest)

		return
	}

	if error := log_entry.AddLogEntry(logEntryRequest.LogEntry); error != nil {
		render.Render(writer, request, handler.ErrorRenderer(error))

		return
	}

	render.Status(request, http.StatusCreated)
	error := render.Render(writer, request, log_entry.NewLogEntryResponse(logEntryRequest.LogEntry))
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}
}

// GET /log-entries/{id}
func GetLogEntry(writer http.ResponseWriter, request *http.Request) {
	logEntryPointer := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)

	error := render.Render(writer, request, log_entry.NewLogEntryResponse(logEntryPointer))
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}
}

// PATCH /log-entries/{id}
func PutLogEntry(writer http.ResponseWriter, request *http.Request) {
	logEntryPointer := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)
	logEntryRequest := &models.LogEntryRequest{}

	if error := render.Bind(request, logEntryRequest); error != nil {
		render.Render(writer, request, handler.ErrBadRequest)

		return
	}

	updatedLogEntryPointer, error := log_entry.UpdateLogEntry((*logEntryPointer).ID, logEntryRequest)
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}

	error = render.Render(writer, request, log_entry.NewLogEntryResponse(updatedLogEntryPointer))
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}
}

// DELETE /log-entries/{id}
func DeleteLogEntry(writer http.ResponseWriter, request *http.Request) {
	logEntry := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)

	error := log_entry.DeleteLogEntry(logEntry.ID)
	if error != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(error))

		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
