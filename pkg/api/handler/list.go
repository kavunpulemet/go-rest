package handler

import (
	"RESTAPIService2/pkg/api/utils"
	"RESTAPIService2/pkg/service/list"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateList(service list.TodoListService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		var input list.TodoList
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := service.Create(userId, input)
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

type getAllListsResponse struct {
	Data []list.TodoList `json:"data"`
}

func GetAllLists(service list.TodoListService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		lists, err := service.GetAll(userId)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := getAllListsResponse{
			Data: lists,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func GetListById(service list.TodoListService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		list, err := service.GetById(userId, id)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(list); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func DeleteList(service list.TodoListService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = service.Delete(userId, id)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func UpdateList(service list.TodoListService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var input list.UpdateListInput
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = service.Update(userId, id, input); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
