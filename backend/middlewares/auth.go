package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowtDev/to-do-app/utils"
)

func Authentication(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	email, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	//setting logged user mail in the context
	context.Set("email", email)
	context.Next()
}