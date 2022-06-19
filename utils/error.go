package utils

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Context string `json:"context,omitempty"`
}

var (
	BadRequestError = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
	InternalServerError = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
	NotFoundError = APIError{
		Code:    http.StatusNotFound,
		Message: "Not found",
	}
)

func WriteError(w http.ResponseWriter, data APIError) {
	j, _ := json.Marshal(data)
	w.WriteHeader(data.Code)
	w.Write(j)
}
