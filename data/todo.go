package data

import (
	"errors"
	"time"
)

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	CreatedBy   string
	CreatedAt   time.Time
}

// GetTodos returns paginated todos
func GetTodos(page, limit int64) ([]Todo, error) {
	var todos []Todo
	offset := (page - 1) * limit
	err := DB.Limit(int(limit)).Offset(int(offset)).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (todo *Todo) CreateTodo(userEmail string) (*Todo, error) {
	todo.CreatedBy = userEmail
	err := DB.Create(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// UpdateTodoById updates a todo if the user is the creator
func (todo *Todo) UpdateTodoById(id uint, userEmail string) (*Todo, error) {
	var existingTodo Todo
	err := DB.First(&existingTodo, id).Error
	if err != nil {
		return nil, err
	}

	if existingTodo.CreatedBy != userEmail {
		return nil, errors.New("unauthorized to update this todo")
	}

	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description

	err = DB.Save(&existingTodo).Error
	if err != nil {
		return nil, err
	}
	return &existingTodo, nil
}

// DeleteTodoById deletes a todo if the user is the creator
func DeleteTodoById(id uint, userEmail string) error {
	var todo Todo
	err := DB.First(&todo, id).Error
	if err != nil {
		return err
	}

	if todo.CreatedBy != userEmail {
		return errors.New("unauthorized to delete this todo")
	}

	return DB.Delete(&todo).Error
}
