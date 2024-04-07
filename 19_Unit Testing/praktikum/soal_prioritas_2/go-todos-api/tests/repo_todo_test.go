package tests

import (
	"github.com/stretchr/testify/mock"
	"go-todos-api/models"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func NewMockRepository() *TodoRepositoryMock {
	return &TodoRepositoryMock{}
}

func (m *TodoRepositoryMock) GetAll() ([]models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) GetByID(id string) (models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) Create(todoInput models.TodoInput) (models.Todo, error) {
	args := m.Called(todoInput)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) Update(todoInput models.TodoInput, id string) (models.Todo, error) {
	args := m.Called(todoInput, id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
