package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyam0507/todo-app/pkg/models"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the todo
	if err := todo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoC, err := todo.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, todoC)
}
