package handler

import (
	"fmt"
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

func ServerErrorMsgRenderer(message string) *ErrorResponse {
	return &ErrorResponse{
		Err:        fmt.Errorf(message),
		StatusCode: 500,
		StatusText: InternalServerErrorStatusText,
		Message:    message,
	}
}
