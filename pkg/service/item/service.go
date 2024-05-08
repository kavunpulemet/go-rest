package item

import (
	todo "RESTAPIService2"
	"RESTAPIService2/pkg/repository"
)

type TodoItemService interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type ImplTodoItem struct {
	repo     repository.TodoItemRepository
	listRepo repository.TodoListRepository
}

func NewTodoItemService(repo repository.TodoItemRepository, listRepo repository.TodoListRepository) *ImplTodoItem {
	return &ImplTodoItem{repo: repo, listRepo: listRepo}
}

func (s *ImplTodoItem) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *ImplTodoItem) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ImplTodoItem) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ImplTodoItem) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ImplTodoItem) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
