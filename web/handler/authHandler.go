package handler

import (
	"back-end/web/model"
	"back-end/web/service"
	"back-end/web/util"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetAuthHandlers() model.Handlers {
	return model.Handlers {
		"/auth/signin": signin,
		"/auth/signup": signup,
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	errRsp, methodErr := validateMethod([]string{"POST"}, r.Method, r.URL.Path)
	if methodErr != nil {
		errRsp.WriteError(w)
		return
	}

	var cfg model.Config
	cfg.GetConfig()

	var creds model.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		model.NewErrorResponse(404, "Invalid json body").WriteError(w)
		return
	}

	var ui *model.User
	ui, err = us.FindUserWithPassword(creds.Username)

	if err != nil || !service.PasswordsMatch(ui.Password, creds.Password)  {
		model.NewErrorResponse(404, "Incorrect password").WriteError(w)
		return
	}

	expirationTime := time.Now().Add(time.Duration(cfg.JwtExpirationTime) * time.Minute)
	claims := &model.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JwtKey))
	if err != nil {
		log.Println(err)
		model.NewErrorResponse(500, "Could not generate token").WriteError(w)
		return
	}

	resp := model.TokenResponse {
		Username: creds.Username,
		Token: tokenString,
	}

	js, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		model.NewErrorResponse(500, "Could not generate token").WriteError(w)
		return
	}
	util.WriteJsonResponse(js, 200, w)
}

func signup(w http.ResponseWriter, r *http.Request) {
	errRsp, methodErr := validateMethod([]string{"POST"}, r.Method, r.URL.Path)
	if methodErr != nil {
		errRsp.WriteError(w)
		return
	}

	var creds model.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		model.NewErrorResponse(400, "Invalid json body").WriteError(w)
		return
	}

	u := model.User {
		Id: 0,
		Username: creds.Username,
		Password: creds.Password,
	}
	
	err = us.CreateUser(u)
	if err != nil {
		model.NewErrorResponse(400, "Username or password already taken").WriteError(w)
		return
	}

	w.WriteHeader(200)
}