package ResultPattern

import (
	"net/http"
)

func NewError(statusCode int, errorCode int, message string, payload *map[string]any) *ErrorBase[map[string]any] {
	if payload == nil {
		payload = &map[string]any{}
	}
	return &ErrorBase[map[string]any]{
		ResultBase: ResultBase[map[string]any]{
			IsSuccess:  false,
			TraceId:    newTraceID(),
			StatusCode: statusCode,
			ErrorCode:  errorCode,
			Message:    message,
			Data:       *payload,
		},
	}
}

// General Context 1000-1099 — transport-agnostic failures

func NewErrorUnAuthorized(message ...string) *ErrorBase[map[string]any] {
	initMessage := "UnAuthorized"
	if len(message) > 0 && message[0] != "" {
		initMessage = message[0]
	}
	return NewError(
		http.StatusUnauthorized,
		1000,
		initMessage,
		&map[string]any{},
	)
}

func NewErrorBadRequest(message string) *ErrorBase[map[string]any] {
	return NewError(
		http.StatusBadRequest,
		1001,
		message,
		&map[string]any{},
	)
}

func NewErrorNotFound(message string) *ErrorBase[map[string]any] {
	return NewError(
		http.StatusNotFound,
		1002,
		message,
		&map[string]any{},
	)
}

func NewErrorConflict(message string) *ErrorBase[map[string]any] {
	return NewError(
		http.StatusConflict,
		1003,
		message,
		&map[string]any{},
	)
}

func NewErrorForbidden(message ...string) *ErrorBase[map[string]any] {
	initMessage := "Forbidden"
	if len(message) > 0 && message[0] != "" {
		initMessage = message[0]
	}
	return NewError(
		http.StatusForbidden,
		1004,
		initMessage,
		&map[string]any{},
	)
}

func NewErrorTooManyRequests(message ...string) *ErrorBase[map[string]any] {
	initMessage := "Too many requests"
	if len(message) > 0 && message[0] != "" {
		initMessage = message[0]
	}
	return NewError(
		http.StatusTooManyRequests,
		1006,
		initMessage,
		&map[string]any{},
	)
}

func NewErrorServiceUnavailable(message ...string) *ErrorBase[map[string]any] {
	initMessage := "Service unavailable"
	if len(message) > 0 && message[0] != "" {
		initMessage = message[0]
	}
	return NewError(
		http.StatusServiceUnavailable,
		1007,
		initMessage,
		&map[string]any{},
	)
}
