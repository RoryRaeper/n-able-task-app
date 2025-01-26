package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RoryRaeper/n-able-task-app/mocks"
	"github.com/RoryRaeper/n-able-task-app/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetTask(t *testing.T) {
	mockService := new(mocks.MockService)
	handler := NewHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks/:id", handler.GetTask)

	task := &models.Task{ID: primitive.NewObjectID(), Title: "Test Task"}
	mockService.On("GetTaskByID", mock.Anything, "1").Return(task, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestListTasks(t *testing.T) {
	mockService := new(mocks.MockService)
	handler := NewHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks", handler.ListTasks)

	tasks := []models.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockService.On("ListTasks", mock.Anything, int64(10), int64(10)).Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks?page=1&offset=10&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestListTasksMalformedRequest(t *testing.T) {
	mockService := new(mocks.MockService)
	handler := NewHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks", handler.ListTasks)

	req, _ := http.NewRequest(http.MethodGet, "/tasks?limit=abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	req, _ = http.NewRequest(http.MethodGet, "/tasks?offset=abc", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	req, _ = http.NewRequest(http.MethodGet, "/tasks?page=abc", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateTask(t *testing.T) {
	mockService := new(mocks.MockService)
	handler := NewHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/tasks", handler.CreateTask)

	task := models.Task{Title: "New Task"}
	createdTask := &models.Task{ID: primitive.NewObjectID(), Title: "New Task"}

	mockService.On("CreateTask", mock.Anything, task).Return(createdTask, nil)

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestUpdateTask(t *testing.T) {
	mockService := new(mocks.MockService)
	handler := NewHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/tasks/:id", handler.UpdateTask)

	task := models.Task{Title: "Updated Task"}
	updatedTask := &models.Task{ID: primitive.NewObjectID(), Title: "Updated Task"}

	mockService.On("UpdateTask", mock.Anything, "1", task).Return(updatedTask, nil)

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteTask(t *testing.T) {
	mockService := new(mocks.MockService)
	handler := NewHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/tasks/:id", handler.DeleteTask)

	mockService.On("DeleteTask", mock.Anything, "1").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
