package tests

import (
	"testing"

	"github.com/gowtDev/to-do-app/data"
)

func TestCreateTodo(t *testing.T) {
	data.InitDB()
	todo := data.Todo{Title: "Test Todo", Description: "Test Description"}
	createdTodo, err := todo.CreateTodo("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if createdTodo.Title != todo.Title {
		t.Errorf("Expected title %s, got %s", todo.Title, createdTodo.Title)
	}
}

func TestGetTodos(t *testing.T) {
	data.InitDB()
	todos, err := data.GetTodos(1, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(todos) == 0 {
		t.Error("Expected at least one todo")
	}
}
