package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowtDev/to-do-app/data"
)

func getTodos(context *gin.Context) {
	page, err := strconv.ParseInt(context.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.ParseInt(context.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos, err := data.GetTodos(page, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todosCount := len(todos)
	context.JSON(http.StatusOK, gin.H{"data": todos, "page": page, "limit": limit, "total": todosCount})
}

func createTodos(context *gin.Context) {
	userEmail := context.GetString("email")

	var todo data.Todo
	err := context.ShouldBind(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoCreated, err := todo.CreateTodo(userEmail)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, todoCreated)
}

func updateTodos(context *gin.Context) {
	userEmail := context.GetString("email")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo data.Todo
	err = context.ShouldBind(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoUpdated, err := todo.UpdateTodoById(uint(id), userEmail)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, todoUpdated)
}

func deleteTodos(context *gin.Context) {
	userEmail := context.GetString("email")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.DeleteTodoById(uint(id), userEmail)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
