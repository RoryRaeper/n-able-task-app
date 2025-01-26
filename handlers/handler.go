package handlers

import (
	"github.com/RoryRaeper/n-able-task-app/services"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service services.Service
}

func NewHandler(service services.Service) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET task",
	})
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET tasks",
	})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST task",
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PUT task",
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE task",
	})
}
