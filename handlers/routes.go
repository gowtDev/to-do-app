package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gowtDev/to-do-app/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	auth := server.Group("/")

	auth.Use(middlewares.Authentication)
	auth.POST("/todos", createTodos)
	auth.PUT("/todos/:id", updateTodos)
	auth.DELETE("/todos/:id", deleteTodos)
	auth.GET("/todos", getTodos)

	//no need token for register and login apis
	server.POST("/register", createUser)
	server.POST("/login", login)
}
