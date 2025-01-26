package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/RoryRaeper/n-able-task-app/models"
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

// Fetches the task with the provided ID
// An error is returned if the task is not found, and invalid ID is provided, or an error occurs while fetching the task
func (h *TaskHandler) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := h.service.GetTaskByID(ctx, id)
	if err != nil {
		log.Printf("Error fetching task %s: %s", id, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
	log.Printf("Task '%s' fetched", id)
}

// Lists all tasks in the database
// The user can provide optional pagination parameters, limit, offset, and page
// If the user provides a page number, the offset is calculated based on the limit and page number
func (h *TaskHandler) ListTasks(ctx *gin.Context) {
	var limit int64 = 100
	var offset int64 = 0
	requestQueries := ctx.Request.URL.Query()

	// Pagination
	limitStr := requestQueries.Get("limit")
	offsetStr := requestQueries.Get("offset")
	pageStr := requestQueries.Get("page")

	if len(limitStr) > 0 {
		queryLimit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit format"})
			return
		}
		limit = queryLimit
	}

	if len(offsetStr) > 0 {
		queryOffset, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset format"})
			return
		}
		offset = queryOffset
	}

	// Calculate offset based on page number
	if len(pageStr) > 0 {
		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page format"})
			return
		}
		offset += +((page - 1) * limit)
	}

	tasks, err := h.service.ListTasks(ctx.Request.Context(), offset, limit)
	fmt.Println(err)
	if err != nil {
		log.Printf("Error fetching tasks: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("offset", strconv.FormatInt(offset, 10))
	ctx.Header("limit", strconv.FormatInt(limit, 10))
	ctx.Header("page", pageStr)
	ctx.JSON(http.StatusOK, tasks)
	log.Println("List tasks call successul")
}

// Creates a task with the provided information
// An error is returned if the json is malformed
// TODO: Add validation for task fields and request body
// Maybe add validation for duplicate titles?
func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	taskRequest := models.Task{}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(body, &taskRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.CreateTask(ctx.Request.Context(), taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
	log.Printf("New task '%s' created", task.ID.Hex())
}

// Updates the task with the provided ID with the updated task information
// Returns an error if the the json is malformed
// TODO: Add validation to ensure the task exists before attempting an update
func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskRequest := models.Task{}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(body, &taskRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := h.service.UpdateTask(ctx.Request.Context(), id, taskRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedTask)
	log.Printf("Task '%s' updated", id)
}

// Deletes the task with the provided ID
func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.DeleteTask(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Task %s deleted successfully", id)})
	log.Printf("Task '%s' deleted", id)
}
