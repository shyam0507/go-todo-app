package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam0507/todo-app/pkg/controllers"
)

var RegisterUserRoutes = func(router *gin.Engine) {
	router.POST("/api/users", controllers.CreateUser)
	router.POST("/api/login", controllers.LoginUser)
}
