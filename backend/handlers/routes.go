package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gowtDev/to-do-app/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	auth := server.Group("/")

	auth.Use(middlewares.Authentication)
	auth.POST("/iam/todos", createTodos)
	auth.PUT("/iam/todos/:id", updateTodos)
	auth.DELETE("/iam/todos/:id", deleteTodos)
	auth.GET("/iam/todos", getTodos)

	//no need token for register and login apis
	server.POST("/iam/register", createUser)
	server.POST("/iam/login", login)
}
