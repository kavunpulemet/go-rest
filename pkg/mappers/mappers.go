package mappers

import (
	"RESTAPIService2/pkg/repository/models"
	"RESTAPIService2/pkg/service/item"
)

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

func MapToUpdateItemInput(UpdateItemInput item.UpdateItemInput) models.UpdateItemInput {
	return models.UpdateItemInput{
		Title:       UpdateItemInput.Title,
		Description: UpdateItemInput.Description,
		Done:        UpdateItemInput.Done,
	}
}
