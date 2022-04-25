package handler

import (
	"back-end/web/model"
	"back-end/web/util"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func GetAuthHandlers() model.Handlers {
	return model.Handlers {
		"/auth/signin": signin,
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

	// TODO add proper db users
	pwd, ok := users[creds.Username]

	if !ok || pwd != creds.Password {
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