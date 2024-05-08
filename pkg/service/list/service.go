package list

import (
	"RESTAPIService2/pkg/mappers"
	"RESTAPIService2/pkg/repository"
)

type TodoListService interface {
	Create(userId int, list TodoList) (int, error)
	GetAll(userId int) ([]TodoList, error)
	GetById(userId, listId int) (TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input UpdateListInput) error
}

type ImplTodoList struct {
	repo repository.TodoListRepository
}

func NewTodoListService(repo repository.TodoListRepository) *ImplTodoList {
	return &ImplTodoList{repo: repo}
}

func (s *ImplTodoList) Create(userId int, list TodoList) (int, error) {
	return s.repo.Create(userId, mappers.MapToTodoList(list))
}

func (s *ImplTodoList) GetAll(userId int) ([]TodoList, error) {
	var todoLists []TodoList

	repositoryTodoLists, err := s.repo.GetAll(userId)
	if err != nil {
		return nil, err
	}

	for _, todoList := range repositoryTodoLists {
		todoLists = append(todoLists, mappers.MapFromTodoList(todoList))
	}

	return todoLists, nil
}

func (s *ImplTodoList) GetById(userId, listId int) (TodoList, error) {
	todoList, err := s.repo.GetById(userId, listId)
	if err != nil {
		return TodoList{}, err
	}

	return mappers.MapFromTodoList(todoList), err
}

func (s *ImplTodoList) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *ImplTodoList) Update(userId, listId int, input UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, mappers.MapToUpdateListInput(input))
}
