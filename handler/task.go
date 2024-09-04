package handler

import (
	"net/http"
	"tracker/model"
	"tracker/repo"

	"github.com/gin-gonic/gin"
)

type TasksHandler struct {
	repo repo.Tasks
}

func NewTasksHandler(repo repo.Tasks) TasksHandler {
	return TasksHandler{repo}
}

func (t TasksHandler) Create(c *gin.Context) {
	var b model.Task
	err := c.ShouldBindJSON(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request!"})
		return
	}

	err = t.repo.Create(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task stored!"})
}
