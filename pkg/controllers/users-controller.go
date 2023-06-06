package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyam0507/todo-app/pkg/models"
	"github.com/shyam0507/todo-app/pkg/utility"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the user
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusCreated, user)
}

func LoginUser(c *gin.Context) {
	var r models.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.ValidateLoginRequest(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.Login()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := utility.GenerateToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusCreated, user)
}
