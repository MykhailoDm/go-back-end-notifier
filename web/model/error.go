package model

import (
	"back-end/web/util"
	"encoding/json"
	"net/http"
	"time"
)

type errorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message string `json:"message"`
	Status int `json:"status"`
}

func NewErrorResponse(st int, msg string) errorResponse {
	return errorResponse{
		Status: st,
		Message: msg,
		Timestamp: time.Now(),
	}
}

func (e errorResponse) WriteError(w http.ResponseWriter) {
	js, _ := json.Marshal(e)
	util.WriteJsonResponse(js, e.Status, w)
}