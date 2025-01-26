package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetTask(c *gin.Context)
	ListTask(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}
