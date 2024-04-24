package item

import (
	"RESTAPIService2/pkg/mappers"
	"RESTAPIService2/pkg/repository"
)

type TodoItemService interface {
	Create(userId, listId int, item TodoItem) (int, error)
	GetAll(userId, listId int) ([]TodoItem, error)
	GetById(userId, itemId int) (TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input UpdateItemInput) error
}

type ImplTodoItem struct {
	repo     repository.TodoItemRepository
	listRepo repository.TodoListRepository
}

func NewTodoItemService(repo repository.TodoItemRepository, listRepo repository.TodoListRepository) *ImplTodoItem {
	return &ImplTodoItem{repo: repo, listRepo: listRepo}
}

func (s *ImplTodoItem) Create(userId, listId int, item TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, mappers.MapToTodoItem(item))
}

func (s *ImplTodoItem) GetAll(userId, listId int) ([]TodoItem, error) {
	var todoItems []TodoItem

	repositoryTodoItems, err := s.repo.GetAll(userId, listId)
	if err != nil {
		return nil, err
	}

	for _, todoItem := range repositoryTodoItems {
		todoItems = append(todoItems, mappers.MapFromTodoItem(todoItem))
	}

	return todoItems, nil
}

func (s *ImplTodoItem) GetById(userId, itemId int) (TodoItem, error) {
	todoItem, err := s.repo.GetById(userId, itemId)
	if err != nil {
		return TodoItem{}, err
	}

	return mappers.MapFromTodoItem(todoItem), nil
}

func (s *ImplTodoItem) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ImplTodoItem) Update(userId, itemId int, input UpdateItemInput) error {
	return s.repo.Update(userId, itemId, mappers.MapToUpdateItemInput(input))
}
