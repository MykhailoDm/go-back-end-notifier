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
		"/status": requestStatus,
	}
}

func requestStatus(w http.ResponseWriter, r *http.Request) {
	errRsp, err := validateMethod([]string{"GET"}, r.Method, r.URL.Path)
	if err != nil {
		errRsp.WriteError(w)
		return
	}

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
		rsp := model.NewErrorResponse(http.StatusInternalServerError, "Issue while building Status response body")
		rsp.WriteError(w)
		return
	}

	util.WriteJsonResponse(js, http.StatusOK, w)
}