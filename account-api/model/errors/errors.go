package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization   Type = "AUTHORIZATION"
	BadRequest      Type = "BADREQUEST"
	Conflict        Type = "CONFLICT"
	Internal        Type = "INTERNAL"
	NotFound        Type = "NOTFOUND"
	PayloadTooLarge Type = "PAYLOADTOOLARGE"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}

	return http.StatusInternalServerError
}

func AuthorizationError(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func BadRequestError(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

func ConflictError(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("Resource: %v with value: %v already exists", name, value),
	}
}

func InternalError() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprint("Internal server error"),
	}
}

func NotFoundError(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("Resource: %v with value: %v not found", name, value),
	}
}

func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
	}
}
