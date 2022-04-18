package main

import (
	"back-end/web/handler"
	"back-end/web/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var cfg model.Config
	cfg.GetConfig()

	handlers := handler.GetStatusHandlers()
	handlers.AddHandlers()

	err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil)
	if err != nil {
		log.Println(err)
	}
}