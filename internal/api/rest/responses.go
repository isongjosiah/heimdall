package rest

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

type ServerResponse struct {
	Err        error       `json:"-"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Payload    interface{} `json:"payload"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
}

// WriteJSONResponse writes the server handler response to the
// response writer
func WriteJSONResponse(rw http.ResponseWriter, statusCode int, content []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if _, err := rw.Write(content); err != nil {
		logger, _ := zap.NewProduction()
		logger.Error("Unable to write json response")
	}
}

// RespondWithError parses an error to the ServerResponse type
func RespondWithError(err error, message, status string, statusCode int) *ServerResponse {
	var wrappedErr error

	if err != nil {
		//wrappedErr = errors.Wrap(err, message)
		wrappedErr = errors.New(err.Error())
	} else {
		wrappedErr = errors.New(message)
	}

	return &ServerResponse{
		Err:        wrappedErr,
		StatusCode: statusCode,
		Message:    message,
		Status:     status,
	}
}

func respondWithError(err error, message string, httpStatusCode int) *ErrorResponse {
	var wrappedErr error
	if err != nil {
		wrappedErr = errors.Wrap(err, message)
	} else {
		wrappedErr = errors.New(message)
	}

	return &ErrorResponse{
		ErrorMessage: wrappedErr.Error(),
		StatusCode:   httpStatusCode,
	}
}

func writeErrorResponse(rw http.ResponseWriter, statusCode int, errString string) {
	r := respondWithError(nil, errString, http.StatusBadRequest)
	errorResponse, _ := json.Marshal(r)
	WriteJSONResponse(rw, statusCode, errorResponse)
}
