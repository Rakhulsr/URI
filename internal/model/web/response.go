package web

import (
	"net/http"

	"github.com/Rakhulsr/go-url-shortener/internal/helper"
)

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(w http.ResponseWriter, code int, status string, msg string, data interface{}) {
	response := WebResponse{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    data,
	}

	err := helper.JsonEncode(w, response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
