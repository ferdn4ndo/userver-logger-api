package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"-"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{
		StatusCode: http.StatusMethodNotAllowed,
		StatusText: "Method not allowed.",
		Message:    "Method not allowed. Check the documentation for the correct endpoints.",
	}
	ErrUnauthorized = &ErrorResponse{
		StatusCode: http.StatusUnauthorized,
		StatusText: "Unauthorized.",
		Message:    "You're not authorized. Please check your credentials and try again.",
	}
	ErrNotFound = &ErrorResponse{
		StatusCode: http.StatusNotFound,
		StatusText: "Resource not found.",
		Message:    "Resource not found. Check the given identifier and try again.",
	}
	ErrBadRequest = &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad request. Check the data sent in the request an try again.",
	}
)

func (errorResponse *ErrorResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	render.Status(request, errorResponse.StatusCode)

	json.NewEncoder(writer).Encode(errorResponse)

	return nil
}

const BadRequestStatusText = "There was a bad request."

func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: BadRequestStatusText,
		Message:    err.Error(),
	}
}

const InternalServerErrorStatusText = "An internal server error has occurred."

func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: InternalServerErrorStatusText,
		Message:    err.Error(),
	}
}

func AddCustomErrorHandlerfunc(writer http.ResponseWriter, request *http.Request, value interface{}) {
	if err, ok := value.(error); ok {

		// We set a default error status response code if one hasn't been set.
		if _, ok := request.Context().Value(render.StatusCtxKey).(int); !ok {
			writer.WriteHeader(400)
		}

		// We log the error
		log.Printf("Logging err: %s\n", err.Error())

		// We change the response to not reveal the actual error message,
		// instead we can transform the message something more friendly or mapped
		// to some code / language, etc.
		render.DefaultResponder(writer, request, render.M{"status": "error"})
		return
	}

	render.DefaultResponder(writer, request, value)
}
