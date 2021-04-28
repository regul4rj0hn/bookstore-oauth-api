package errors

import (
	"errors"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    string `json:"data"`
}

func Message(msg string) error {
	return errors.New(msg)
}

func BadRequest(data string) *Response {
	return &Response{
		Message: "bad_request",
		Status:  http.StatusBadRequest,
		Data:    data,
	}
}

func Conflict(data string) *Response {
	return &Response{
		Message: "conflict",
		Status:  http.StatusConflict,
		Data:    data,
	}
}

func InternalServerError(data string) *Response {
	return &Response{
		Message: "internal_server_error",
		Status:  http.StatusInternalServerError,
		Data:    data,
	}
}

func NotFound(data string) *Response {
	return &Response{
		Message: "not_found",
		Status:  http.StatusNotFound,
		Data:    data,
	}
}
