package service

import (
	"back-end/web/model"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
	models model.Models
}

var lock = &sync.Mutex{}
var us *UserService

func GetUserService(m model.Models) *UserService {
    if us == nil {
        lock.Lock()
        defer lock.Unlock()
        if us == nil {
            fmt.Println("Creating user service instance")
            us = &UserService{
				models: m,
			}
        }
    }

    return us
}

func(us *UserService) ValidateJwt(w http.ResponseWriter, r *http.Request) (*model.UserInfo, error) {
	claims := &model.Claims{}

	cfg := model.Config{}
	cfg.GetConfig()

	authHeader := r.Header.Get("Authorization")
	if len(authHeader) <= 7 {
		model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized: invalid authorization header").WriteError(w)
		return nil, errors.New("Unathorized")
	}
	
	bearerToken := authHeader[7:]
	tkn, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JwtKey), nil
	})

	if err != nil || !tkn.Valid {
		model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized: invalid token").WriteError(w)
		return nil, errors.New("Unathorized")
	}
	
	var ui *model.UserInfo
	ui, err = us.FindUser(claims.Username)
	if err != nil {
		model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized: user not found").WriteError(w)
		return nil, errors.New("Unathorized")
	}
	if ui == nil || ui.Id == 0 || ui.Username == "" {
		model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized: user not found").WriteError(w)
		return nil, errors.New("Unathorized")
	}

	return ui, nil
}

func (us *UserService) CreateUser(user model.User) error {
	log.Printf("Creating user: %v", user.Username)
	
	var hpwd string
	hpwd, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hpwd
	err = us.models.DB.CreateUser(user)
	return err
}

func (us *UserService) FindUser(username string) (*model.UserInfo, error) {
	log.Printf("Retrieving user: %v", username)
	ui, err := us.models.DB.GetUser(username)
	return ui, err
}

func (us *UserService) FindUserWithPassword(username string) (*model.User, error) {
	log.Printf("Retrieving user: %v", username)
	u, err := us.models.DB.GetUserWithPassword(username)
	return u, err
}

func hashPassword(pwd string) (string, error) {
	var passwordBytes = []byte(pwd)
  
	hpwdBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
  
	return string(hpwdBytes), err
  }
  
func PasswordsMatch(hpwd, cpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(cpwd))
	return err == nil
}