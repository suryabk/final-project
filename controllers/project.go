package controllers

import (
	"final-project/database"
	"final-project/repository"
	"final-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProject(c *gin.Context) {
	var (
		result gin.H
	)

	project, err := repository.GetAllProject(database.DbConnection)

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

func InsertProject(c *gin.Context) {
	var input structs.InputProject

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	project := structs.Project{
		ProjectName: input.ProjectName,
		Description: input.Description,
		Budget:      input.Budget,
		// Deadline:    input.Deadline,
		CreatedBy: userID,
	}

	err = repository.InsertProject(database.DbConnection, project)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Project",
	})
}

func UpdateProject(c *gin.Context) {
	var input structs.InputProject
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	project := structs.Project{
		ProjectName: input.ProjectName,
		Description: input.Description,
		Budget:      input.Budget,
		// Deadline:    input.Deadline,
		CreatedBy: userID,
	}

	project.ProjectID = int(id)
	err = repository.UpdateProject(database.DbConnection, project)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Project",
	})
}

func DeleteProject(c *gin.Context) {
	var project structs.Project
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Project ID"})
		return
	}

	project.ProjectID = id

	err = repository.DeleteProject(database.DbConnection, project)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Project",
	})
}
