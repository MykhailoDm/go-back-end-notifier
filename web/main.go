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

	db, err := model.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler.GetStatusHandlers().AddHandlers()
	handler.GetAuthHandlers().AddHandlers()

	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil)
	if err != nil {
		log.Println(err)
	}
}