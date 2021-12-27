package http

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func responseWriter(w http.ResponseWriter, responseCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(data)
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	res := HTTPResponse{
		Success: true,
		Data:    data,
		Errors:  []string{},
	}
	responseWriter(w, http.StatusOK, res)
}

func ResponseBadRequest(w http.ResponseWriter, messages []string) {
	res := HTTPResponse{
		Success: false,
		Data:    nil,
		Errors:  messages,
	}
	responseWriter(w, http.StatusBadRequest, res)
}

func ResponseFatal(w http.ResponseWriter, messages []string) {
	res := HTTPResponse{
		Success: false,
		Data:    nil,
		Errors:  messages,
	}
	responseWriter(w, http.StatusInternalServerError, res)
}
