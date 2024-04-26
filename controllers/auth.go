package controllers

import (
	"final-project/database"
	"final-project/middleware"
	"final-project/repository"
	"final-project/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var user structs.User
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.RegisterUser(database.DbConnection, user)

	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
			return
		} else {
			panic(err)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Register Success",
	})
}

func LoginHandler(c *gin.Context) {

	var (
		result gin.H
		user   structs.User
	)
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	loginRes, err := repository.LoginUser(database.DbConnection, user)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else if loginRes.UserID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not registered"})
		return
	} else if loginRes.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	} else {
		loginInfo := structs.LoginReturn{UserID: loginRes.UserID, Email: loginRes.Email}
		result = gin.H{
			"result": loginInfo,
		}
	}

	token := middleware.GenerateToken(loginRes.UserID, loginRes.Email)
	c.Header("Authorization", "Bearer "+token)
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, result)
}
