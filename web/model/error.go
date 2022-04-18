package model

import (
	"back-end/web/util"
	"encoding/json"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message string `json:"message"`
	Status int `json:"status"`
}

func NewErrorResponse(st int, msg string) ErrorResponse {
	return ErrorResponse{
		Status: st,
		Message: msg,
		Timestamp: time.Now(),
	}
}

func (e ErrorResponse) WriteError(w http.ResponseWriter) {
	js, _ := json.Marshal(e)
	util.WriteJsonResponse(js, e.Status, w)
}