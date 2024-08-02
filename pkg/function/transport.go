package function

import (
	"heimdall/internal/value"
	"net/http"
)

func StatusCode(status string) int {
	switch status {
	case value.Success:
		return http.StatusOK
	case value.NotFound:
		return http.StatusNotFound
	case value.Created:
		return http.StatusCreated
	default:
		return http.StatusInternalServerError
	}
}
