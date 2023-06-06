package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam0507/todo-app/pkg/controllers"
	"github.com/shyam0507/todo-app/pkg/middleware"
)

var RegisterTodoRoutes = func(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())
	router.POST("/api/todo", controllers.CreateTodo)

}
