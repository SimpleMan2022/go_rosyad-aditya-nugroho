package tests

import (
	"errors"
	"go-todos-api/models"
	"go-todos-api/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTodoService_GetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expectedTodos := []models.Todo{{
			ID:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   gorm.DeletedAt{},
			Title:       "contoh",
			Description: "deskripsi",
		}}
		todoRepository.On("GetAll").Return(expectedTodos, nil)
		todos, err := todoServices.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, expectedTodos, todos, "Result must be todos")
		todoRepository.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expected := []models.Todo{
			{},
		}
		expectedError := errors.New("Todos is empty")
		todoRepository.On("GetAll").Return(expected, expectedError)
		todos, err := todoServices.GetAll()
		assert.NotNil(t, err)
		assert.Equal(t, expected, todos)
		assert.EqualError(t, err, expectedError.Error())
		todoRepository.AssertExpectations(t)
	})
}

func TestTodoService_GetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expectedTodo := models.Todo{
			ID:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   gorm.DeletedAt{},
			Title:       "Ngoding",
			Description: "Alta",
		}
		todoRepository.On("GetByID", "1").Return(expectedTodo, nil)
		todo, err := todoServices.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, expectedTodo, todo)
		todoRepository.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expectedTodo := models.Todo{}
		expectedError := errors.New("Todo not Found")
		todoRepository.On("GetByID", "1").Return(expectedTodo, expectedError)
		todo, err := todoServices.GetByID("1")
		assert.Error(t, err)
		assert.Equal(t, expectedTodo, todo)
		assert.EqualError(t, err, expectedError.Error())
		todoRepository.AssertExpectations(t)
	})
}

func TestTodoService_Create(t *testing.T) {
	todoInput := models.TodoInput{
		Title:       "desain",
		Description: "desain ui",
	}
	t.Run("Success", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)

		expectedTodo := models.Todo{
			ID:          2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   gorm.DeletedAt{},
			Title:       "desain",
			Description: "desain ui",
		}
		todoRepository.On("Create", todoInput).Return(expectedTodo, nil)
		newTodo, err := todoServices.Create(todoInput)

		assert.NoError(t, err)
		assert.Equal(t, expectedTodo.Title, newTodo.Title)
		todoRepository.AssertExpectations(t)
	})
	t.Run("Failed", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expectedError := errors.New("Create new todo failed")
		todoRepository.On("Create", todoInput).Return(models.Todo{}, expectedError)
		newTodo, err := todoServices.Create(todoInput)

		assert.Error(t, err)
		assert.Equal(t, models.Todo{}, newTodo)
		assert.EqualError(t, err, expectedError.Error())
		todoRepository.AssertExpectations(t)
	})
}

func TestTodoService_Update(t *testing.T) {
	todoInput := models.TodoInput{
		Title:       "ngoding",
		Description: "ngoding golang",
	}
	mockTodo := models.Todo{
		ID:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Test Todo",
		Description: "Test Description",
	}
	t.Run("Success", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		todoRepository.On("GetByID", "1").Return(mockTodo, nil)
		mockTodo.Title = todoInput.Title
		mockTodo.Description = todoInput.Description
		todoRepository.On("Update", todoInput, "1").Return(mockTodo, nil)
		updatedTodo, err := todoServices.Update(todoInput, "1")

		assert.NoError(t, err)
		assert.Equal(t, mockTodo.Title, updatedTodo.Title)
	})
	t.Run("Failed Get Id", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expectedError := errors.New("Todo not found")
		todoRepository.On("GetByID", "5").Return(models.Todo{}, expectedError)
		todo, err := todoServices.GetByID("5")
		assert.Error(t, err)
		assert.Equal(t, models.Todo{}, todo)
		assert.EqualError(t, err, expectedError.Error())
		todoRepository.AssertExpectations(t)
	})
	t.Run("Failed Update", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		todoRepository.On("GetByID", "1").Return(mockTodo, nil)
		mockTodo.Title = todoInput.Title
		mockTodo.Description = todoInput.Description
		expectedError := errors.New("Update todo failed")
		todoRepository.On("Update", todoInput, "1").Return(models.Todo{}, expectedError)
		updatedTodo, err := todoServices.Update(todoInput, "1")

		assert.Error(t, err)
		assert.Equal(t, models.Todo{}, updatedTodo)
		assert.EqualError(t, err, expectedError.Error())
	})
}

func TestTodoService_Delete(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		todoRepository.On("Delete", "1").Return(nil)
		err := todoServices.Delete("1")
		assert.NoError(t, err)
		todoRepository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		todoRepository := NewMockRepository()
		todoServices := services.NewTodoService(todoRepository)
		expectedError := errors.New("Delete todo failed")
		todoRepository.On("Delete", "1").Return(expectedError)
		err := todoServices.Delete("1")
		assert.Error(t, err)
		assert.EqualError(t, err, expectedError.Error())
		todoRepository.AssertExpectations(t)
	})

}
