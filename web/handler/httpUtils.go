package handler

import (
	"back-end/web/model"
	"log"
	"net/http"
	"errors"
)

func validateMethod(methods []string, rm string, p string) (model.ErrorResponse, error) {
	for _, m := range methods {
		if m == rm {
			log.Printf("Request. Method: %v. Path: %v.", rm, p)
			return model.ErrorResponse{}, nil
		}
	}

	log.Printf("Request Method Not Allowed. Method: %v. Path: %v", rm, p)
	return model.NewErrorResponse(http.StatusMethodNotAllowed, "Method not allowed."), errors.New("invalid method")
}