package util

import "net/http"

func WriteJsonResponse(js []byte, st int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}