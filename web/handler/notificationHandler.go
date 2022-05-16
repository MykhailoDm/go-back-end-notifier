package handler

import (
	"back-end/web/model"
	"back-end/web/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetNotificationHandlers() model.Handlers {
	return model.Handlers{
		"/api/v1/notifications": notificationsHandlerFunc,
		"/api/v1/notifications/": notificationsHandlerFuncById,
	}
}

func notificationsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	errRsp, methodErr := validateMethod([]string{"GET", "POST"}, r.Method, r.URL.Path)
	if methodErr != nil {
		errRsp.WriteError(w)
		return
	}

	ui, err := us.ValidateJwt(w, r)
	if err != nil {
		return
	}

	switch r.Method {
	case "GET":
		getNotifications(ui.Id, w, r)
	case "POST":
		createNotification(ui.Id, w, r)
	}
}

func notificationsHandlerFuncById(w http.ResponseWriter, r *http.Request) {
	errRsp, methodErr := validateMethod([]string{"GET", "DELETE", "PUT"}, r.Method, r.URL.Path)
	if methodErr != nil {
		errRsp.WriteError(w)
		return
	}

	_, err := us.ValidateJwt(w, r)
	if err != nil {
		return
	}

	lastSlashIndex := strings.LastIndex(r.URL.Path, "/")
	nidString := r.URL.Path[lastSlashIndex+1:]
	nid, err := strconv.Atoi(nidString)
	if err != nil {
		model.NewErrorResponse(400, "Invalid id format").WriteError(w)
		return
	}

	switch r.Method {
	case "GET":
		getNotification(nid, w, r)
	case "DELETE":
		deleteNotification(nid, w, r)
	case "PUT":
		updateNotification(nid, w, r)
	}
}

func getNotifications(id int, w http.ResponseWriter, r *http.Request) {
	notifications, err := ns.GetNotifications(id)
	if err != nil {
		model.NewErrorResponse(400, "Could not retrieve notifications for this user").WriteError(w)
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

func getNotification(id int, w http.ResponseWriter, r *http.Request) {
	n, err := ns.GetNotification(id)
	if err != nil {
		model.NewErrorResponse(400, "Could not retrieve notification for this user").WriteError(w)
		return
	}

	js, err := json.Marshal(n)
	if err != nil {
		log.Println(err)
		model.NewErrorResponse(500, "Could not generate notifications response").WriteError(w)
		return
	}
	util.WriteJsonResponse(js, 200, w)
}

func createNotification(uid int, w http.ResponseWriter, r *http.Request) {
	var n model.Notification
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		model.NewErrorResponse(400, "Invalid json body").WriteError(w)
		return
	}

	n.UserId = uid
	err = ns.CreateNotification(n)
	if err != nil {
		model.NewErrorResponse(400, "Bad Request").WriteError(w)
		return
	}
}

func deleteNotification(id int, w http.ResponseWriter, r *http.Request) {
	err := ns.DeleteNotification(id)
	if err != nil {
		model.NewErrorResponse(404, "Not Found").WriteError(w)
		return
	}
}

func updateNotification(id int, w http.ResponseWriter, r *http.Request) {
	var n model.Notification
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		model.NewErrorResponse(400, "Invalid json body").WriteError(w)
		return
	}

	err = ns.UpdateNotification(id, n)
	if err != nil {
		model.NewErrorResponse(404, "Not Found").WriteError(w)
		return
	}
}