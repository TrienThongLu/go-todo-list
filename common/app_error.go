package common

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, root error, message, log, key string) *AppError {
	return &AppError{statusCode, root, message, log, key}
}

func NewErrorResponse(root error, message, log, key string) *AppError {
	return &AppError{http.StatusBadRequest, root, message, log, key}
}
