package handler

import (
	todo "RESTAPIService2"
	"RESTAPIService2/pkg/api/utils"
	"RESTAPIService2/pkg/service/auth"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SignUp(service auth.AuthorizationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input todo.User

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := service.CreateUser(input)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"id": id,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(service auth.AuthorizationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input signInInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		token, err := service.GenerateToken(input.Username, input.Password)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"token": token,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			logrus.Println("Error encoding JSON response:", err)
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
