package errors

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	StatusCode int    `json:"-"`
	StatusText string `json:"status"`
	Message    string `json:"message"`
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func Message(msg string) error {
	return errors.New(msg)
}

func BadRequest(msg string) *Response {
	return &Response{
		StatusText: "bad_request",
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func Conflict(msg string) *Response {
	return &Response{
		StatusText: "conflict",
		StatusCode: http.StatusConflict,
		Message:    msg,
	}
}

func InternalServerError(msg string) *Response {
	return &Response{
		StatusText: "internal_server_error",
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}

func NotFound(msg string) *Response {
	return &Response{
		StatusText: "not_found",
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}
