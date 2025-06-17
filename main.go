package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gowtDev/to-do-app/data"
	"github.com/gowtDev/to-do-app/handlers"
)

func main() {
	data.InitDB()
	server := gin.Default()
	handlers.RegisterRoutes(server)
	server.Run(":8080")
}
