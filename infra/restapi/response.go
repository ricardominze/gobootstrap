package restapi

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, message string, code int, errors interface{}) {
	response := Response{
		Status:  "error",
		Message: message,
		Code:    code,
		Errors:  errors,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
