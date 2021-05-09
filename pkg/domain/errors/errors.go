package errors

import (
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrBadRequest     = &Response{StatusCode: http.StatusBadRequest, StatusText: "bad_request"}
	ErrNotFound       = &Response{StatusCode: http.StatusNotFound, StatusText: "not_found"}
	ErrConflict       = &Response{StatusCode: http.StatusConflict, StatusText: "conflict"}
	ErrInternalServer = &Response{StatusCode: http.StatusInternalServerError, StatusText: "internal_server_error"}
)

type Response struct {
	StatusCode int    `json:"-"`
	StatusText string `json:"status"`
	Message    string `json:"message"`
}

func BadRequest(msg string) *Response {
	return &Response{
		StatusText: ErrBadRequest.StatusText,
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NotFound(msg string) *Response {
	return &Response{
		StatusText: ErrNotFound.StatusText,
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func Conflict(msg string) *Response {
	return &Response{
		StatusText: ErrConflict.StatusText,
		StatusCode: http.StatusConflict,
		Message:    msg,
	}
}

func InternalServerError(msg string) *Response {
	return &Response{
		StatusText: ErrInternalServer.StatusText,
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
