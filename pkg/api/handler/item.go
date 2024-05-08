package handler

import (
	todo "RESTAPIService2"
	"RESTAPIService2/pkg/api/utils"
	"RESTAPIService2/pkg/service/item"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateItem(service item.TodoItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		listId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var input todo.TodoItem
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := service.Create(userId, listId, input)
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

func GetAllItems(service item.TodoItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		listId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		items, err := service.GetAll(userId, listId)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(items); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func GetItemById(service item.TodoItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		itemId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		item, err := service.GetById(userId, itemId)
		if err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(item); err != nil {
			utils.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func DeleteItem(service item.TodoItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		itemId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = service.Delete(userId, itemId)
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

func UpdateItem(service item.TodoItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		itemId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var input todo.UpdateItemInput
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = service.Update(userId, itemId, input); err != nil {
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
