package model

import "net/http"

type Handlers map[string]func(w http.ResponseWriter, r *http.Request)

func (h Handlers) AddHandlers() {
	for path, handler := range h {
		http.HandleFunc(path, handler)
	}
}