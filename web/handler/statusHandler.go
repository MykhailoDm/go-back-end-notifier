package handler

import (
	"back-end/web/model"
	"back-end/web/util"
	"encoding/json"
	"log"
	"net/http"
)

func GetStatusHandlers() model.Handlers {
	return model.Handlers {
		"/status": getStatus,
	}
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	var cfg model.Config
	cfg.GetConfig()

	currentStatus := model.AppStatus {
		Status: "Available",
		Environment: cfg.Env,
		Version: cfg.Version,
	}

	js, err := json.Marshal(currentStatus)
	if err != nil {
		log.Println(err)
	}

	util.WriteJsonResponse(js, http.StatusOK, w)
}