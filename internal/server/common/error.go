package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	ErrorCode   string `json:"error_code"`
}

func NotFound(w http.ResponseWriter) {
	JSONError(w, http.StatusNotFound, ErrorResponse{
		Code:        10000,
		Description: "Unknown request",
		ErrorCode:   "CF-NotFound",
	})
}

func JSONError(w http.ResponseWriter, status int, errorResponse ErrorResponse) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}
