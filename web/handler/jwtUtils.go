package handler

import (
	"back-end/web/model"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)


func ValidateJwt(w http.ResponseWriter, r *http.Request) error {
	claims := &model.Claims{}

	cfg := model.Config{}
	cfg.GetConfig()

	tkn, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtKey), nil
	})

	if err != nil || !tkn.Valid {
		model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized").WriteError(w)
		return errors.New("Unathorized")
	}
	return nil
}