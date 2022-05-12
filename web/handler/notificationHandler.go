package handler

import (
	"back-end/web/model"
	"back-end/web/util"
	"encoding/json"
	"log"
	"net/http"
)

func GetNotificationHandlers() model.Handlers {
	return model.Handlers{
		"/api/1/notifications": notificationsHandlerFunc,
	}
}

func notificationsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	validateMethod([]string{"GET", "POST",}, r.Method, r.URL.Path)

	ui, err := us.ValidateJwt(w, r)
	if err != nil {
		return
	}

	switch r.Method {
	case "GET": getNotifications(ui.Id, w, r)
	case "POST":	
	}
}

func getNotifications(id int, w http.ResponseWriter, r *http.Request) {
	notifications, err := ns.GetNotifications(id)
	if err != nil {
		model.NewErrorResponse(404, "Invalid json body").WriteError(w)
		return
	}

	js, err := json.Marshal(notifications)
	if err != nil {
		log.Println(err)
		model.NewErrorResponse(500, "Could not generate notifications response").WriteError(w)
		return
	}
	util.WriteJsonResponse(js, 200, w)
}
