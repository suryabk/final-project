package controllers

import (
	"final-project/database"
	"final-project/repository"
	"final-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTask(c *gin.Context) {
	var (
		result gin.H
	)

	project, err := repository.GetAllTask(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": project,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertTask(c *gin.Context) {
	var input structs.Task

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	err = repository.InsertTask(database.DbConnection, input)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Task",
	})
}

func UpdateTask(c *gin.Context) {
	var input structs.Task
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	input.TaskID = int(id)
	err = repository.UpdateTask(database.DbConnection, input)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Task",
	})
}

func DeleteTask(c *gin.Context) {
	var task structs.Task
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	task.TaskID = id

	err = repository.DeleteTask(database.DbConnection, task)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Task",
	})
}
