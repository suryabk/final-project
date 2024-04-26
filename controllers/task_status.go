package controllers

import (
	"final-project/database"
	"final-project/repository"
	"final-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTaskStatus(c *gin.Context) {
	var (
		result gin.H
	)

	status, err := repository.GetAllTaskStatus(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": status,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertTaskStatus(c *gin.Context) {
	var input structs.TaskStatus

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	err = repository.InsertTaskStatus(database.DbConnection, input)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Task Status",
	})
}

func UpdateTaskStatus(c *gin.Context) {
	var input structs.TaskStatus
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	input.StatusID = int(id)
	err = repository.UpdateTaskStatus(database.DbConnection, input)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Task Status",
	})
}

func DeleteTaskStatus(c *gin.Context) {
	var status structs.TaskStatus
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task Status ID"})
		return
	}

	status.StatusID = id

	err = repository.DeleteTaskStatus(database.DbConnection, status)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Task Status",
	})
}
