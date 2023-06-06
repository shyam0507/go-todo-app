package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyam0507/todo-app/pkg/utility"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := utility.ValidateToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			fmt.Println(err)
			context.Abort()
			return
		}
		context.Next()
	}
}
