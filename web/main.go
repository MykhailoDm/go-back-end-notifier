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

	m := model.Models {
		DB: model.DBModel{
			DB: db,
		},
	}
	handler.LoadServices(m)

	handler.GetStatusHandlers().AddHandlers()
	handler.GetAuthHandlers().AddHandlers()
	handler.GetNotificationHandlers().AddHandlers()

	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil)
	if err != nil {
		log.Println(err)
	}
}