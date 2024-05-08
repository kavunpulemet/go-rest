package mappers

import (
	"RESTAPIService2/pkg/repository/models"
	"RESTAPIService2/pkg/service/item"
	"RESTAPIService2/pkg/service/list"
)

func MapToTodoList(todoList list.TodoList) models.TodoList {
	return models.TodoList{
		Id:          todoList.Id,
		Title:       todoList.Title,
		Description: todoList.Description,
	}
}

func MapFromTodoList(todoList models.TodoList) list.TodoList {
	return list.TodoList{
		Id:          todoList.Id,
		Title:       todoList.Title,
		Description: todoList.Description,
	}
}

func MapToUpdateListInput(updateListInput list.UpdateListInput) models.UpdateListInput {
	return models.UpdateListInput{
		Title:       updateListInput.Title,
		Description: updateListInput.Description,
	}
}

func MapToTodoItem(todoItem item.TodoItem) models.TodoItem {
	return models.TodoItem{
		Id:          todoItem.Id,
		Title:       todoItem.Title,
		Description: todoItem.Description,
		Done:        todoItem.Done,
	}
}

func MapFromTodoItem(todoItem models.TodoItem) item.TodoItem {
	return item.TodoItem{
		Id:          todoItem.Id,
		Title:       todoItem.Title,
		Description: todoItem.Description,
		Done:        todoItem.Done,
	}
}

func MapToUpdateItemInput(updateItemInput item.UpdateItemInput) models.UpdateItemInput {
	return models.UpdateItemInput{
		Title:       updateItemInput.Title,
		Description: updateItemInput.Description,
		Done:        updateItemInput.Done,
	}
}
